# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-11-03

First stable release.

### API Stability Notice

This is the first stable release (v1.0.0). All public APIs are considered stable and will maintain backward compatibility within the v1.x series. Breaking changes will only occur in v2.0.0+ releases, following Semantic Versioning.

**Stable APIs include:**
- Money creation: `New()`, `NewZero()`, `MustNew()`
- Arithmetic: `Add()`, `Subtract()`, `Multiply()`, `Divide()`
- Allocation: `Allocate()`, `AllocateByPercentage()`
- Comparisons, utilities, rounding, serialization, and currency functions
- All error variables and RoundScheme constants

---

## [Unreleased]

Future features and improvements will be documented here.