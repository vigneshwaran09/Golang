# Example
> Let's imagine we have a function like this.
![alt text](img1.png)

> When we invoke the function which will return fucntion as a value.
![alt text](img2.png)

> So you might be wondering how is this debitFunc function accessing this amount
![alt text](img3.png)

> Like you expect that once this activate gift card function returns and finish
![alt text](img4.png)

> This variable amount which's part of this activate gift card function scope would be erased from memory right
 ![alt text](img5.png)

> So how does debitFunc still have the amount to even reference or subtract from well the answer to that my friends is closure
 ![alt text](img6.png)

> The first time this function gets called we're setting an intial amount of 100 
 ![alt text](img7.png)

> And then we create a debitFunc
 ![alt text](img8.png)

> That can later use the gift card and subtract from the intial amount
 ![alt text](img9.png)
 ![alt text](img7.png)

> Now since the debitFunc function depends on a variable outside of itself 
 ![alt text](img10.png)

> Under the hood the programming language will Enclose both the debitFunc function and any variable that depends on that are defined outside of itself
 ![alt text](img11.png)

> And store it  in  memory as an enclosed unit or closure
 ![alt text](img12.png)

> So what's being referenced in memory by this use gift card one variable is actually something that looks like this and for use gift Card 2 we're once again calling activate gift card a second time and when we call it the second time another separate closure will get stored in memory and referenced by use gift card 2
 ![alt text](img13.png)

> So when we call both of these use gift card functions they're actually subtracting from their own amounts that are contained within own individual closures in memory.
 ![alt text](img14.png)

<hr>


[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/jHd0FczIjAE/0.jpg)](https://www.youtube.com/watch?v=jHd0FczIjAE)


