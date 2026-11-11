// The CUDA version of Primes.java: the same naive prime test, one thread per element.
//
// Build and run on a machine with an NVIDIA GPU:
//   nvcc -O2 -o primes primes.cu && ./primes 10000

#include <cstdio>
#include <cstdlib>

// START1 OMIT
__device__ bool isPrime(int n) {
    if (n < 2) return false;
    for (int d = 2; d < n; d++) {
        if (n % d == 0) return false;
    }
    return true;
}

// The body of the Parallel For: one thread per element
//   thread i:  0        1        2        ...
//   numbers:   1000003  1000003  1000003  ...
//   result:    true     true     true     ...
__global__ void primes(const int *numbers, bool *result, int count) {
    int i = blockIdx.x * blockDim.x + threadIdx.x;
    if (i >= count) return;
    result[i] = isPrime(numbers[i]);
}
// END1 OMIT

static bool allTrue(const bool *values, int count) {
    for (int i = 0; i < count; i++) {
        if (!values[i]) return false;
    }
    return true;
}

int main(int argc, char **argv) {
    int count = argc > 1 ? atoi(argv[1]) : 10000;

    // START2 OMIT
    int *numbers;
    bool *result;
    cudaMallocManaged(&numbers, count * sizeof(int)); // visible to CPU and GPU
    cudaMallocManaged(&result, count * sizeof(bool));
    for (int i = 0; i < count; i++) numbers[i] = 1000003; // a prime number

    int threads = 256;                               // threads per block
    int blocks = (count + threads - 1) / threads;    // enough blocks for all elements
    primes<<<blocks, threads>>>(numbers, result, count);
    cudaDeviceSynchronize();                         // wait for the GPU
    // END2 OMIT

    cudaEvent_t start, stop;
    cudaEventCreate(&start);
    cudaEventCreate(&stop);
    cudaEventRecord(start);
    primes<<<blocks, threads>>>(numbers, result, count);
    cudaEventRecord(stop);
    cudaEventSynchronize(stop);
    float ms = 0;
    cudaEventElapsedTime(&ms, start, stop);

    cudaError_t err = cudaGetLastError();
    if (err != cudaSuccess) {
        fprintf(stderr, "CUDA error: %s\n", cudaGetErrorString(err));
        return 1;
    }
    printf("%d numbers, all prime: %s, %.0f ms\n", count, allTrue(result, count) ? "true" : "false", ms);

    cudaFree(numbers);
    cudaFree(result);
    return 0;
}
