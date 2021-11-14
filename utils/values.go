package main

import "fmt"

func byval(q *int) {
	fmt.Printf("3. passed--integer(q)	%T:	*q=i=%v			&q=%p		q=&i=%p	\n", q, *q, &q, q	)
	*q = 4143
	fmt.Printf("4. set--integer(*q)	%T:	*q=i=%v		&q=%p		q=&i=%p			\n", q, *q, &q, q	)
	q = nil
}

func main() {
	fmt.Printf("                	%%T(type)	%%v			%%p			%%p					\n"					)
	i := int(42)
	fmt.Printf("1. main--integer	%T:	i=%v 			&i=%p						\n", i, i, &i		)
	p := &i
	fmt.Printf("2. main--pointer	%T:	*p=i=%v  	&p=%p		p=&i=%p				\n", p, &p, *p, p	)
	byval(p)
	fmt.Printf("5. main--pointer	%T:	*p=i=%v	&p=%p	p=&i=%p						\n", p, &p, *p, p	)
	fmt.Printf("6. main--integer	%T:	i=%v 		&i=%p							\n", i, &i, i		)
}
