# Comparable
- comparable means we can compare two structs using equal operator,if all field of the struct only in primitive.
- But any of the field is in dynamic memory like `slice` etc so we can't compare by using `==` operator.