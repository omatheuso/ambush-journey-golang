package main

import "fmt"

func stringifyAge(value: IPerson): string {
    return `${value.age} years old`;
}

interface IPerson {
    age: number
}

interface ICar {
    age: number
    model: string
}

const car: ICar = {
    age: 28,
    model: `Ford`,
}

stringifyAge(car); // OK

func main() {
	stringifyAge(car);
}
