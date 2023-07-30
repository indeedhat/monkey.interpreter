# Monkey lang interpreter

Implementation of monkey lang based on [this book](https://interpreterbook.com/) with a bunch of my own changes + features.

Some or all of the features i have added may be a later part of the book, dont know, haven't gotten there yet.

> current place 151

## Current state of monkey lang (to the eval stage)

### String literals
```
"This is a string";
"This is a string with \"escaped\" quates"
```

### Integer Literals
```
12345;
-1;
```

### Boolean Literals
```
true;
false;
```

### prefix operators
```
!true;
-1;
```

### infix operators
```
true == true;
true != false;
0 < 1;
1 > 0;
1 >= 0;
0 <= 1;
1 + 1;
1 - 1;
1 * 2;
4 / 2;
```

### if/else statements
```
if (1 > 5) {
    false
} else if (1 == 5) {
    false
} else {
    true
}
```

## Misc
```
return 1;
null;
```

## Funcitions/closures
```
let add = fn(a, b) {
    a + b;
};

let call_twice = fn(f) {
    f();
    f();
};

call_twice(fn() {
    1 + 2;
});
```
