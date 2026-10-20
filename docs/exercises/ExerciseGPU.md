# Exercise - GPU Programming with Shadertoy

If you do not finish during the lecture period, please finish it as homework.

All tasks run in the browser on <https://www.shadertoy.com>. You don't need to install anything, and you can run shaders without an account. You need an account only to save them, so copy your code into a local file from time to time.

If Shadertoy is not reachable, use the editor of [The Book of Shaders](https://thebookofshaders.com) or [GLSL Sandbox](https://glslsandbox.com) instead.

## Explore

- Click through some shaders on Shadertoy, e.g. the most popular ones. Open the code of at least two of them.
- Pick one shader and explain to your neighbor what it does. Which part of the code is responsible for which part of the image?
- Look at the FPS counter below the image. Which shaders are fast, which ones are slow? Can you guess why?

## Task 1 - Prime numbers

Time: about 25 minutes.

In the lecture, we tested prime numbers on the CPU and the GPU. Now you do the same in Shadertoy: every pixel stands for one number and shows whether this number is prime.

### Step 0

Create a new shader (<https://www.shadertoy.com/new>). Show a horizontal gradient:

```glsl
void mainImage(out vec4 fragColor, in vec2 fragCoord) {
    fragColor = vec4(vec3(fragCoord.x / iResolution.x), 1.0);
}
```

Before you run it: What will the image look like?

### From 2D to 1D

Shadertoy is a 2D Parallel For. Every pixel gets its coordinate `(x, y)`. But for the prime numbers you need **one** number per pixel.

- How do you get a 1D index from the 2D pixel coordinates, so that every pixel gets its own number?
- Check your index: Divide it by the number of pixels and show it as a color, like in Step 0. What do you expect to see if your index is correct?

### The prime test

- Write a function `bool isPrime(int n)`.
- Show prime numbers white and all other numbers black.
- Look at the image. Do you see a pattern?
- Add a large offset to your index, e.g. test the numbers from 100000000 onwards. What happens to the FPS? Why?

## Task 2 - Mandelbrot

The Mandelbrot set is computed with a loop for every pixel. Start with `z = 0` and repeat

    z = z² + c

where `c` is the coordinate of the pixel as a complex number. If `|z|` grows larger than 2, the point is not in the set. Stop after a maximum number of iterations, e.g. 200.

Hint: A complex number fits into a `vec2`. The square of `(x, y)` is `(x² − y², 2xy)`.

- Before you run it: Predict which areas of the image need many iterations and which areas need only a few.
- Show the number of iterations as a color. This is a map of how much work every pixel does.
- Animate it with `iTime`, e.g. zoom into the image.
- Which areas are expensive, and why?

## Common mistakes

- `float f = 1;` does not compile. GLSL has no implicit conversion from `int` to `float`. Write `1.0` or `float(i)`.
- `fragCoord` contains the center of the pixel: `(0.5, 0.5)`, `(1.5, 0.5)`, ... Use `int(fragCoord.x)` or `ivec2(fragCoord)` if you need whole numbers.
- The origin `(0, 0)` is at the **bottom left**, not at the top left.
- Recursion is not allowed in GLSL.

## If you finish early

- Make something nice: a plasma effect with `sin()` of the coordinates and `iTime`, or animated shapes with distance functions (see the circle from the lecture).
- [compute.toys](https://compute.toys): Shadertoy for compute shaders (WebGPU). Try the prime numbers as a real compute shader. Needs a browser with WebGPU support.
- [GPU-Puzzles](https://github.com/srush/GPU-Puzzles): Learn the CUDA programming model with small puzzles in Python (Google Colab).

## After this Exercise

- You can write simple shaders in GLSL
- You understand the GPU as a Parallel For, where you only write the body
- You can explain why loops with different lengths are expensive on the GPU
