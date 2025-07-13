# 当前是场景 3

## 总览

- [场景 1](https://github.com/relax-space/gomod1A)
- [场景 2](https://github.com/relax-space/gomod2A)
- [场景 3](https://github.com/relax-space/gomod3A)

A 依赖 B, B 依赖 C

场景 1: A 没有直接依赖了 C, A 通过 B 调用 C 或者(A 虽然没有通过 B 调用 C,但是 A 调用 B 的 M 方法,M 方法所在的包中有别的方法调用 C), 是否 go mod tidy 可以自动添加, 需要手动处理吗?

可以, 并且会把 B 对 C 的依赖也添加进来,显示 indirect,

场景 2: A 直接依赖了 C, 但是 A 没有通过 B 调用 C,是否 go mod tidy 可以自动添加, 需要手动处理吗?

可以, 不会把 B 对 C 的依赖也添加进来

场景 3: A 直接依赖了 C, 同时 A 通过 B 调用 C,是否 go mod tidy 可以自动添加, 需要手动处理吗? 如果体现有冲突 以及如何处理

可以, 会把 B 对 C 的依赖也添加进来,版本不一样, 就都依赖进来就好
