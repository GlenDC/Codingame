+++++++++++ number of digits to output
> number 1
+ initial number
>>>> number 5
++++++++++++++++++++++++++++++++++++++++++++ (comma)
> number 6
++++++++++++++++++++++++++++++++ (space)
<<<<<< number 0
[
  > number 1
  copy number 1 to number 7
  [>>>>>>+>+<<<<<<<-]>>>>>>>[<<<<<<<+>>>>>>>-]

  <
  divide number 7 by 10 (begins in number 7)
  [
    >
    ++++++++++  set the divisor number 8
    [
      subtract from the dividend and divisor
      -<-
      if dividend reaches zero break out
        copy dividend to number 9
        [>>+>+<<<-]>>>[<<<+>>>-]
        set number 10
        +
        if number 9 clear number 10
        <[>[-]<[-]]
        if number 10 move remaining divisor to number 11
        >[<<[>>>+<<<-]>>[-]]
      jump back to number 8 (divisor possition)
      <<
    ]
    if number 11 is empty (no remainder) increment the quotient number 12
    >>> number 11
    copy to number 13
    [>>+>+<<<-]>>>[<<<+>>>-]
    set number 14
    +
    if number 13 clear number 14
    <[>[-]<[-]]
    if number 14 increment quotient
    >[<<+>>[-]]
    <<<<<<< number 7
  ]

  quotient is in number 12 and remainder is in number 11
  >>>>> number 12
  if number 12 output value plus offset to ascii 0
  [++++++++++++++++++++++++++++++++++++++++++++++++.[-]]
  subtract number 11 from 10
  ++++++++++  number 12 is now 10
  < number 11
  [->-<]
  > number 12
  output number 12 even if it's zero
  ++++++++++++++++++++++++++++++++++++++++++++++++.[-]
  <<<<<<<<<<< number 1

  check for final number
  copy number 0 to number 3
  <[>>>+>+<<<<-]>>>>[<<<<+>>>>-]
  <- number 3
  if number 3 output (comma) and (space)
  [>>.>.<<<[-]]
  << number 1

  [>>+>+<<<-]>>>[<<<+>>>-]<<[<+>-]>[<+>-]<<<-
]