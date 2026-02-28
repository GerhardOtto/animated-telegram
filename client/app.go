package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gerhardotto/animated-telegram/client/backendservice"
	"github.com/gerhardotto/animated-telegram/client/internal/config"
	"github.com/gerhardotto/animated-telegram/client/internal/service"
	"github.com/gerhardotto/animated-telegram/client/internal/sorting"
)

// App wires together configuration and the gRPC connection.
type App struct {
	conn     *grpc.ClientConn
	cfg      *config.Config
	greeting *service.GreetingService
	auth     *service.AuthService
	data     *service.DataService
}

// NewApp parses configuration and dials the gRPC server.
func NewApp() (*App, error) {
	cfg := config.Parse()

	conn, err := grpc.NewClient(
		cfg.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(timingInterceptor),
	)
	if err != nil {
		return nil, err
	}

	client := pb.NewDataBackendClient(conn)

	return &App{
		conn:     conn,
		cfg:      cfg,
		greeting: service.NewGreetingService(client),
		auth:     service.NewAuthService(client),
		data:     service.NewDataService(client),
	}, nil
}

// Close releases the underlying gRPC connection.
func (a *App) Close() {
	a.conn.Close()
}

// Run executes the client workflow.
func (a *App) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := a.greeting.SayHello(ctx, a.cfg.Name); err != nil {
		return err
	}

	token, err := a.auth.GetToken(ctx, a.cfg.Name)
	if err != nil {
		return err
	}

	sections, err := a.data.GetTypesOfData(ctx, a.cfg.Name, token)
	if err != nil {
		return err
	}

	allData, err := a.data.GetAllData(ctx, a.cfg.Name, token, sections)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(os.Stdin)
	for section, items := range allData {
		fmt.Printf("\nDataset: %s (%d items)\n", section, len(items))

		var algo, order string
		if a.cfg.CustomSort {
			algo = prompt(reader, "Sort algorithm (merge/quick/insertion)", a.cfg.Sort)
			order = prompt(reader, "Sort order (asc/desc)", a.cfg.Order)
		} else {
			algo = a.cfg.Sort
			order = a.cfg.Order
		}

		cmp := comparatorFor(items, order)
		sorted := applySort(algo, items, cmp)

		printItems(section, sorted)
	}
	return nil
}

// prompt prints a label with a default value and reads a line from stdin.
// If the user enters nothing, the default is returned.
func prompt(reader *bufio.Reader, label, defaultVal string) string {
	fmt.Printf("%s [%s]: ", label, defaultVal)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	if line == "" {
		return defaultVal
	}
	return line
}

// comparatorFor returns the appropriate CompareFunc by inspecting the first item:
// if its stringval is non-empty the dataset is string-typed, otherwise integer-typed.
func comparatorFor(items []*pb.DataItem, order string) sorting.CompareFunc {
	byString := len(items) > 0 && items[0].GetStringval() != ""
	if byString {
		if order == "desc" {
			return sorting.ByStringValDesc
		}
		return sorting.ByStringValAsc
	}
	if order == "desc" {
		return sorting.ByIntValDesc
	}
	return sorting.ByIntValAsc
}

// applySort dispatches to the chosen algorithm, falling back to MergeSort on unknown input.
func applySort(algo string, items []*pb.DataItem, cmp sorting.CompareFunc) []*pb.DataItem {
	switch strings.ToLower(algo) {
	case "quick":
		return sorting.QuickSort(items, cmp)
	case "insertion":
		return sorting.InsertionSort(items, cmp)
	default:
		return sorting.MergeSort(items, cmp)
	}
}

// printItems prints a preview of up to 10 sorted items from a dataset.
func printItems(section string, items []*pb.DataItem) {
	limit := len(items)
	if limit > 10 {
		limit = 10
	}
	fmt.Printf("Sorted %s (showing %d/%d):\n", section, limit, len(items))
	for i := 0; i < limit; i++ {
		item := items[i]
		fmt.Printf("  [%d] intVal=%d stringval=%q\n", i+1, item.GetIntVal(), item.GetStringval())
	}
}
