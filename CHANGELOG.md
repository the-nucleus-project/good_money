# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Refactored upcoming features documentation in README

### Fixed
- Improved project structure: moved package from root to `goodmoney/` directory
- Updated installation instructions and import paths

---

## [2.0.0] - 2025-11-16

### Changed
- **BREAKING**: Changed package name from `good_money` to `goodmoney` for better Go idiomatic naming
- Updated module path to `github.com/the-nucleus-project/goodmoney`
- Updated all import paths and documentation

### Fixed
- Fixed import path references in README and examples

---

## [1.0.0] - 2025-11-03

### Added
- Initial stable release
- Money type with precise arithmetic using `int64` minor units
- Full ISO 4217 currency support (180+ currencies)
- 8 rounding schemes (HalfUp, HalfDown, HalfEven, etc.)
- Money allocation methods: `Allocate()` and `AllocateByPercentage()`
- Overflow/underflow protection
- JSON serialization support
- Database driver support (`Scan`/`Value`)
- Comprehensive test coverage

### API Stability Notice

This is the first stable release (v1.0.0). All public APIs are considered stable and will maintain backward compatibility within the v1.x series. Breaking changes will only occur in v2.0.0+ releases, following Semantic Versioning.

**Stable APIs include:**
- Money creation: `New()`, `NewZero()`, `MustNew()`
- Arithmetic: `Add()`, `Subtract()`, `Multiply()`, `Divide()`
- Allocation: `Allocate()`, `AllocateByPercentage()`
- Comparisons, utilities, rounding, serialization, and currency functions
- All error variables and RoundScheme constants