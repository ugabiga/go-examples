# DI Examples

# Thoughts on DI Libraries

## 1. GoIoC
I don't like how it works because I have to specify the target by giving a string name for the injection.

## 2. Dig
Overall, it was good. I like that you just need to set the code to inject, and removing the library seems easy, except for `dig.In` However, my biggest complaint is the ambiguity of the errors when debugging.

## 3. Fx
I'd choose `fx` over `dig` because, although it has a bit more boilerplate, it makes debugging in `dig` less uncomfortable.
