# GO_BlackJack

* Author: Kyle Shultz, Stephanie Ball, Adam Torek, Cameron Sewell
* Class: CS 354
* Semester: Fall 2022

## Overview

This is a blackjack game that was implemented with Go. It is meant to be run locally where the player can choose to play with two other AI players and an AI dealer. 

## Reflection

This project was a little trickier than we were were thinking it would be.  Initially, we all thought that Blackjack would be a relatively easy to implement game with it's simple options. However, as we did more research into Blackjack and the various strategies that pros employ we realized that there was more complexity to Blackjack than we inititally thought there would be.

There were two sections of the project that were hard to complete.  One was the AI design and implementation. In general, most of the AI wasn't too bad as we went with a rule based design. The aggresiveness was determinted by a co-factor that would just increase the dealer's tendency to hit and push their score closer to 21. The difficult part was trying to figure out how to implement blackjack actions like split. As much research as we did we could never really determine when it was best for a pro player to split their hand and consequently didn't include that action in the AI logic. 

The other difficult section was dynamically upadating the cards on the game screen. We ended up using a couple of helper methods that allowed us to turn a card object into a card image.  We also used a helper method that took the Hand object and a pointer to the container we wanted to change, the container being a horizontal box of card images that represented the player or the dealers hand. With both of these we could take the cards in the hand and convert them into card images and then take that array of card images and place them into the container we pointed to. This method was called in various parts in the game logic that required the cards on the screen to update. 

If we had more time we probably would have implemented a custom layout using Fyne for our GUI. We would have built out more actions for our AI like split. Ultimately, we would have loved to move this to being a web based game and we think Go would've been the perfect language to implement an online game for. We just did not have the time required to scale the project in these ways but the basic logic is there and moving it to be web based would not take as much work as creating the game did. 

We all enjoyed working with Go and would work with it again.  

## Compiling and Using

Make sure to download Go and gcc (for the GUI) if you haven't already and link both in the path variable of whatever OS you're using. 

We recommend cloning this repository using this command in your terminal:

git clone git@github.com:CameronSewell/GO_BlackJack.git

Once you have the repository pulled down we recommend using an IDE such as VSCode or GoLand. 

If you are using VSCode you will want to install the recommended Go extensions for the best experience. 

You can then run the program from the main.go file using the IDE's run command. From there the GUI should guide you through a game of Blackjack. 
