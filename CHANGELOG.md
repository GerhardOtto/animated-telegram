# Changelog

## [1.1.0](https://github.com/GerhardOtto/animated-telegram/compare/v1.0.0...v1.1.0) (2026-02-28)


### Features

* Add GitHub Actions workflow for running tests ([2a35a94](https://github.com/GerhardOtto/animated-telegram/commit/2a35a94b9629df12a1c7627090b4601a7289eca3))
* Add unit tests for Auth ([2a35a94](https://github.com/GerhardOtto/animated-telegram/commit/2a35a94b9629df12a1c7627090b4601a7289eca3))
* Add unit tests for Data ([2a35a94](https://github.com/GerhardOtto/animated-telegram/commit/2a35a94b9629df12a1c7627090b4601a7289eca3))
* Add unit tests for Greeting ([2a35a94](https://github.com/GerhardOtto/animated-telegram/commit/2a35a94b9629df12a1c7627090b4601a7289eca3))
* Add unit tests for Sorting ([2a35a94](https://github.com/GerhardOtto/animated-telegram/commit/2a35a94b9629df12a1c7627090b4601a7289eca3))
* **service:** Add tests and workflow ([#10](https://github.com/GerhardOtto/animated-telegram/issues/10)) ([2a35a94](https://github.com/GerhardOtto/animated-telegram/commit/2a35a94b9629df12a1c7627090b4601a7289eca3))


### Bug Fixes

* **service:** Remove broken chunk size retry logic in GetAllData ([#12](https://github.com/GerhardOtto/animated-telegram/issues/12)) ([ef0469d](https://github.com/GerhardOtto/animated-telegram/commit/ef0469d3223fe021c92a56eaf0889e342f3d3215))
* **workflow:** Skip tests on release-please PRs and cache go.sum ([#13](https://github.com/GerhardOtto/animated-telegram/issues/13)) ([ae0e6aa](https://github.com/GerhardOtto/animated-telegram/commit/ae0e6aa176892c3b6084167a389b1cf93cbe8e3d))

## 1.0.0 (2026-02-28)


### Features

* Add auth endpoint ([#7](https://github.com/GerhardOtto/animated-telegram/issues/7)) ([c3d4ed4](https://github.com/GerhardOtto/animated-telegram/commit/c3d4ed42cf984a9eff05eb1ff19405561c7de03a))
* Add GetAllData endpoint ([7176a2d](https://github.com/GerhardOtto/animated-telegram/commit/7176a2ded3931966f704a29ec1dcd267aa182b70))
* Add GetTypesOfData endpoint ([7176a2d](https://github.com/GerhardOtto/animated-telegram/commit/7176a2ded3931966f704a29ec1dcd267aa182b70))
* Add Go gRPC client scaffold ([7719ab2](https://github.com/GerhardOtto/animated-telegram/commit/7719ab2a4d2a81c2b9d0864448b1fefd1a3f983a))
* Add greeting endpoint ([#5](https://github.com/GerhardOtto/animated-telegram/issues/5)) ([8977e76](https://github.com/GerhardOtto/animated-telegram/commit/8977e76c9811505668d0735a208e519f490b5cce))
* Add InsertionSort ([adc7a15](https://github.com/GerhardOtto/animated-telegram/commit/adc7a155f8e61100672ed02488ffa54ff8867030))
* Add MergeSort ([adc7a15](https://github.com/GerhardOtto/animated-telegram/commit/adc7a155f8e61100672ed02488ffa54ff8867030))
* Add QuickSort ([adc7a15](https://github.com/GerhardOtto/animated-telegram/commit/adc7a155f8e61100672ed02488ffa54ff8867030))
* Add README ([#3](https://github.com/GerhardOtto/animated-telegram/issues/3)) ([d2ed4b7](https://github.com/GerhardOtto/animated-telegram/commit/d2ed4b752b0e22d1759dd175f13cbfd8369a6a18))
* Add Release-Please ([#1](https://github.com/GerhardOtto/animated-telegram/issues/1)) ([1b586ed](https://github.com/GerhardOtto/animated-telegram/commit/1b586edff7c712daa14a26e96b5e858b29be979c))
* Add script to generate Go client from proto file ([7719ab2](https://github.com/GerhardOtto/animated-telegram/commit/7719ab2a4d2a81c2b9d0864448b1fefd1a3f983a))
* Add timing interceptor ([#6](https://github.com/GerhardOtto/animated-telegram/issues/6)) ([5b94830](https://github.com/GerhardOtto/animated-telegram/commit/5b948305ec48fef4ca755d9febe6d900a0daf4ec))
* Fetch all data ([#8](https://github.com/GerhardOtto/animated-telegram/issues/8)) ([7176a2d](https://github.com/GerhardOtto/animated-telegram/commit/7176a2ded3931966f704a29ec1dcd267aa182b70))
* Implement interactive sorting ([adc7a15](https://github.com/GerhardOtto/animated-telegram/commit/adc7a155f8e61100672ed02488ffa54ff8867030))
* Implement sorting algorithms  ([#9](https://github.com/GerhardOtto/animated-telegram/issues/9)) ([adc7a15](https://github.com/GerhardOtto/animated-telegram/commit/adc7a155f8e61100672ed02488ffa54ff8867030))
* Scaffold grpc client ([#4](https://github.com/GerhardOtto/animated-telegram/issues/4)) ([7719ab2](https://github.com/GerhardOtto/animated-telegram/commit/7719ab2a4d2a81c2b9d0864448b1fefd1a3f983a))


### Bug Fixes

* Implement DataService ([7176a2d](https://github.com/GerhardOtto/animated-telegram/commit/7176a2ded3931966f704a29ec1dcd267aa182b70))
