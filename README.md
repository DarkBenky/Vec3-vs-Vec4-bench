# Benchmarking `Vec3` vs. `Vec4` Performance and the Impact of `math32.Sqrt`

## Overview

This project benchmarks two vector structs, `Vec3` (3 `float32` components) and `Vec4` (4 `float32` components), across vector operations: addition, subtraction, dot product, normalization, and cross product. The study highlights inefficiencies caused by poor mathematical implementations, especially in `math32.Sqrt`.

## Benchmark Results

| Benchmark                     | Iterations  | Time (ns/op) |
|-------------------------------|-------------|--------------|
| **Addition**                  |             |              |
| `Vec3Add`                     | 1,000,000,000 | 0.2539     |
| `Vec4Add`                     | 1,000,000,000 | 0.2541     |
| **Subtraction**               |             |              |
| `Vec3Sub`                     | 1,000,000,000 | 0.2558     |
| `Vec4Sub`                     | 1,000,000,000 | 0.2558     |
| **Dot Product**               |             |              |
| `Vec3Dot`                     | 1,000,000,000 | 0.2552     |
| `Vec4Dot`                     | 1,000,000,000 | 0.2574     |
| **Normalization**             |             |              |
| `Vec3Normalize`               | 1,000,000,000 | 0.2532     |
| `Vec4Normalize`               | 1,000,000,000 | 0.2544     |
| `Vec4NormalizeMath32`         | 206,433,376   | 5.696      |
| **Cross Product**             |             |              |
| `Vec3Cross`                   | 1,000,000,000 | 0.2539     |
| `Vec4Cross`                   | 1,000,000,000 | 0.2550     |

## Key Findings

1. **Performance Parity Between `Vec3` and `Vec4`**  
   Most operations have nearly identical performance due to efficient CPU handling of `float32`.

2. **Normalization Bottleneck with `math32.Sqrt`**  
   - `Vec4NormalizeMath32` is significantly slower (5.696 ns/op).
   - `math32.Sqrt` introduces overhead compared to the standard `math.Sqrt`.

3. **Poor `math32.Sqrt` Implementation**  
   - Uses inefficient or iterative methods.
   - Lacks the hardware-optimized precision of `math.Sqrt`.

## Recommendations

1. **Avoid `math32.Sqrt` for Critical Applications**  
   Use `math.Sqrt` with type casting for better performance.

2. **Optimize `math32.Sqrt`**  
   Implement SIMD or hardware-accelerated methods if constrained to `math32`.

3. **Profile and Validate**  
   - Use profiling tools to detect bottlenecks.
   - Benchmark on various CPUs to ensure consistency.

## Conclusion

Efficient mathematical operations are crucial in performance-critical applications. While `Vec3` and `Vec4` show negligible performance differences, reliance on suboptimal libraries like `math32.Sqrt` can significantly degrade performance. Careful library selection and optimization are key to achieving high-performance results.
