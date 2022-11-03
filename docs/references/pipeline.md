# Pipeline language

The following is the Pipeline processor language definition. With the gradual support of different syntaxes, this document will be adjusted, added or deleted to varying degrees. The following is the Pipeline processor language definition. With the gradual support of different syntaxes, this document will be adjusted, added or deleted to varying degrees.

## Identifiers and Keywords

### Identifier

Identifiers are used to identify objects and can be used to represent a variable, function, etc. The identifier contains keywords

Custom identifiers cannot duplicate keywords in the Pipeline data processor language

Identifiers can consist of numbers (`0-9`), letters (`A-Z a-z`), and underscores (`_`), but the first character cannot be a number and is case sensitive:

- `_abc`
- `abc`
- `abc1`
- `abc_1_`

If you need to start with a letter or use backticks outside of the above characters in an identifier:

- `` `1abc` ``
- `` `@some-variable` ``
- `` `This is an emoji variable👍` ``

### special identifier

The special identifier (`_`) represents the external raw input data when the ppl script is triggered, this parameter may be implicitly passed to some functions

In some functions, `_` will be treated as an alias for `message` for forward compatibility.

### Keywords

Keywords are words with special meaning, such as `if`, `elif`, `else`, `for`, `in`, `break`, `continue`, etc.

## comments

Use `#` as a line comment character, inline comments are not supported

```python
# This is a line comment
a = 1 # this is a line comment

"""
This is a (multi-line) string that replaces the comment
"""
a = 2

"string"
a = 3
```

## Built-in data types

In the Pipeline data processor language, the type of the value of a variable can change dynamically, but each value has its data type, which can be one of the **basic types**, or a **composite type**.

### basic type

#### Integer (int) type

Integer type length is 64bit, signed, currently only supports writing integer literals in decimal, such as `-1`, `0`, `1`, `+19`

#### floating point (float) type

The type of floating-point type is 64bit in length and signed. Currently, only floating-point literals can be written in decimal, such as `-1.00001`, `0.0`, `1.0`, `+19.0`

#### boolean type

There are only two types of boolean literals, `true` and `false`

#### character (str) string type

String literals can be written in double or single quotes, and multi-line strings can be written using triple double or triple single quotes to enclose the content

  * `"hello world"`

  * `'hello world'`

  * ```
    """hello
    world"""
    ```

  * ```
    '''
    hello
    world
    '''
    ```

#### Nil type
   nil is a special data type that means no value. The literal of this type is `nil`. When a variable is used without assignment, its value is nil

### Compound Types

The map type is different from the list type and the basic type. Multiple variables can point to the same map or list object. When assigning a value, the memory copy of the list or map is not performed, but a reference is made.

#### Map type

The map type is a key-value structure, only the string type can be used as the key, and the data type of the value is not limited

It can read and write elements in the map through index expressions

```python
a = {
  "1": [1, "2", 3, nil],
  "2": 1.1,
  "abc": nil,
  "def": true
}

# Since a["1"] is a list object，b just refers to the value of a["1"]
b = a["1"]

"""
now a["1"][0] == 1.1
"""
b[0] = 1.1
```

#### List type

The list type can store any number of values of any type in a list
It can read and write elements in the list through index expressions

```python
a = [1, "2", 3.0, false, nil, {"a": 1}]

a = a[0] # a == 1
```

## operator

The following operators are currently supported by PPL. The higher the value, the higher the priority.

|Priority|Symbol|Associativity|Description|
|-|-|-|-|
| 1 | `=` | right | assignment; named arguments; lowest precedence |
| 2 | `||` | Left | Logical OR |
| 3 | `&&` | left | logical AND |
| 4 | `>=` | left | condition "greater than or equal to" |
| 4 | `>` | left | condition "greater than" |
| 4 | `!=` | left | condition "not equal to" |
| 4 | `==` | left | condition "equals" |
| 4 | `<=` | left | condition "less than or equal to" |
| 4 | `<` | left | condition "less than" |
| 5 | `+` | left | arithmetic "plus" |
| 5 | `-` | left | arithmetic "subtract" |
| 6 | `*` | left | arithmetic "multiply" |
| 6 | `/` | left | arithmetic "divide" |
| 6 | `%` | left | arithmetic "remainder"|
| 7 | `[]` | left | use list subscript or map key value|
| 7 | `()` | None | Can change operator precedence; function call |

## expression

PPL uses the symbol comma `,` as the expression separator, such as parameter passing for calling expressions and the separation of expressions during initialization of map and list

### call expression

The following is a function call to get the number of list elements:

```txt
len([1, 3, "5"])
```

### Binary expressions

An assignment expression is a binary expression that has a return value

````txt
# 0
2 / 5

# 0.4, promote the type of the left operand to a floating-point number when calculating
2 / 5.0

# true
1 + 2 * 3 == 7 && 1 <= 2

# b == 3;
# Due to the right associativity of the `=` operator, a = (b = 3), a == 3
a = b = 3
````

### List initializer

```txt
[1, true, "1", nil]
```

### Map initializer

```txt
{
  "a": 1,
  "b": "2",
}
```

### bracket expression

Bracket expressions can change operand operation precedence in binary expressions, but not associativity

```txt
# 1 + 2 * 3 == 7

(1 + 2) * 3  # == 9
```

## statement

All expressions in the PPL can be regarded as value statements. When the expression ends with the statement delimiter `;` or `\n`, it will be regarded as a statement. For example, the following script contains four statements

```go
len("abc")
1
a = 2; a + 2 * 3 % 2
```

### Value Statement (Expression Statement)

When an expression is followed by a statement separator, it can be regarded as a value statement. The following are four legal statements

```txt
a = 1; b = 2;
d = [1, 2]
len(d)
```

### select statement

The PPL supports the `if/elif/else` syntax:

```txt
if condition {

}
```

```txt
if condition {

} else {

}
```

```txt
if condition_1 {

} elif condition_2 {

} ... elif condition_n {

} else {

}
```

Like most programming languages, according to whether the conditions of `if/elif` are true, enter the corresponding statement block, if not, enter the else branch.

The current condition can be any expression, as long as its value is one of the built-in data types, the following are the predicate conditions:

* When the condition is a value of type `int`, it is `0` then the condition is `false`, otherwise it is `true`

* When the condition is a value of type `float`, it is `0.0` then the condition is `false`, otherwise it is `true`

* When the condition is a value of type `string`, it is an empty string `""` then the condition is `false`, otherwise it is `true`

* When the condition is a value of type `bool`, the condition is the current value

* When the condition is a value of type `nil`, the condition is `false`

* When the condition is a value of type `map`, its length is 0 and the condition is `false`, otherwise it is `true`

* When the condition is a value of type `list`, its length is 0 and the condition is `false`, otherwise it is `true`

### loop statement

PPL supports `for` statement and `for in` statement

The following are two statements that are only allowed in loop blocks:

- `cotinue` statement, do not execute subsequent statements, continue to start the next loop
- `break` statement, which ends the loop

Use of `for` statement may cause infinite loop, should be used with caution, or use `for in` statement instead

```txt
for init-expr; condition; loop-expr {

}
```

```txt
for varb_name in map_value/list_value/string_value  {

}
```

Example of use:

1. Use `for` to execute 10 loops

```txt
for a = 0; a < 10; a = a + 1 {
  
}
```

2. Use `for in` to iterate over all elements of a list

```txt
b = "2"
for a in ["1", "a" ,"2"] {
  b = b + a
  if b == "21a" {
    break
  }
}

# b == "21a"
```

3. Use `for in` to iterate over all the keys of the map

```txt
d = 0
map_a = {"a": 1, "b":2}
for x in map_a {
  d = d + map_a[x]
}
```

4. Use `for in` to iterate over all characters of string

```txt
s = ""
for c in "abcdef" {
  if s == "abc" {
    break
  } else {
    continue
  }
  s = s + "a"
}

# s == "abc"
```
