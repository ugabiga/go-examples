# DI Examples

# Thoughts on DI Libraries

## 1. GoIoC
I don't like how it works because I have to specify the target by giving a string name for the injection.

## 2. Dig
Overall, it was good. I like that you just need to set the code to inject, and removing the library seems easy, 
except for `dig.In` However, my biggest complaint is the ambiguity of the errors when debugging.

## 3. Fx
I prefer `fx` over `dig`. `fx` reduces the amount of repetitive code I had to write in `dig`.
Also debugging is much easier in `fx` than in `dig`.

## 4. Wire
I don’t like that `wire` uses code generation.
And I don't like the fact that if I want to initialize an object, 
I have to list all its dependencies in the `wire.Build` function.
Unless I specifically need compile-time dependency injection, I would not use it personally.