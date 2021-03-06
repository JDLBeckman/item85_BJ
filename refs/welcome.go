package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    var _ error
    var gender string
    var name string
    fmt.Print("Hi there! Welcome to the world of Pokémon! My name is Professor Juniper. Everyone calls me the Pokémon Professor!\nThat's right! This world is widely inhabited by mysterious creatures called Pokémon! Pokémon have mysterious powers. They come in many shapes and live in many different places. We humans live happily with Pokémon! Living and working together, we complement each other. We help each other to accomplish difficult tasks. Having Pokémon battle one another is particularly popular, and it deepens the bonds between people and Pokémon. And that is why I research Pokémon. Well, that's enough from me\nCould you tell me about yourself? Are you a boy? Or a girl? ")
    gender, _ = reader.ReadString('\n')
    gender = gender[:len(gender)-1]
    fmt.Printf("\nYou're a %s, right?\nI'd like to know your name. Please tell me. ", gender)
    name, _ = reader.ReadString('\n')
    name = name[:len(name)-1]
    fmt.Printf("\nSo your name's %s. What a wonderful name! Well then. I'm going to introduce you to your two best friends!\n", name)
    fmt.Println("This young man is Cheren. He can be a little difficult, but he's a very honest person. This young woman is Bianca. She's a little flighty, but she works very hard. I think you three have potential, so I'm going to give you a very, very important Pokémon.")
    fmt.Printf("%s! The moment you choose the Pokémon that will accompany you on your journey, your story will truly begin. During your journey, you will meet many Pokémon and people with different personalities and points of view! I really hope you find what is important to you in all of these travels… That's right! Befriend new people and Pokémon and grow as a person! That is the most important goal for your journey! Let's go visit the world of Pokémon!\n", name)

}
