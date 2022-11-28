# Composition in Golang

This example is inspired from [THIS](https://www.geeksforgeeks.org/composition-in-golang/) article. 


Composition is a method employed to <i><u>write re-usable segments</u></i> of code. It is achieved when objects are made up of other smaller objects with particular behaviors, in other words, Larger objects with a wider functionality are embedded with smaller objects with specific behaviors. The end goal of composition is the <i><u>same as that of inheritance</u></i>, however, instead of inheriting features from a parent class/object, larger objects in this regard are composed of the other objects and can thereby use their functionality.

Since writing effective code in the Go Language revolves majorly around using structs and interfaces, composition is a key takeaway from what the language has to offer. In this article we will discuss the use of the composition of types with structs and interfaces, through type embedding.

