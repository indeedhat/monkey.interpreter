# Monkey lang interpreter

Implementation of monkey lang based on [this book](https://interpreterbook.com/) with a bunch of my own changes + features.

Some or all of the features i have added may be a later part of the book, dont know, haven't gotten there yet.

> current place 177

## Current state of monkey lang (to the eval stage)

### Comments
```
// single line comments
/*
    block comments
*/

/* closing block comments is optional
let commented = "out code";
```

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

### arrays
```
let a = [1];
let a = append(a, 2);
let b = append(a, 3, 4, 5);

len(a) == 2;
len(b) == 5;
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
"Hello," + " " + "World!"
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

### Print
```
print("print me to the screen");

let a = [0, 1, 2];
printf("the array has %d elements", len(a));

let str
```
