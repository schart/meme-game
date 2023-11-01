**Endpoints:**
    root endpoint: game -> /game/?

    POST -> game/room/create-room -> For create a room
        requirements body: nil

    POST -> game/room/join-room -> For join a room
        requirements body: RoomId

    GET ->  game/items/rooms -> For get all rooms
        requirements token: Access with access key

    Web Socket start with GET ->  game/start -> For start game
        requirements url parameter: room link


**Create room:**
    Before creating the room, we check if the player has the token.
        We get the token from cookies, we check.

    Based on the parameters expected from the client, the parameters of the previously created data structure are compared with the incoming parameters.
        Get the HTTP request body.
        Compare with []interface{}.

    We get the account id that we store in the received token and check for the existence of the player.
        If the player exists, continue the process.
        If not, pause and declare "that the player cannot be found."

    We have the id of the player and with this we will check if the player belongs to a room or not.
        If not, proceed.
        If he/she is, pause and declare "that he/she can only join 1 room at a time."

    Create the room and then join the player to its room.

**Join room:**
    Before setting up the room, check if the player has the token.
        We get the token from cookies, we check.

    Based on the parameters expected from the client, the parameters of the previously created data structure are compared with the incoming parameters.
        Get the HTTP request body.
        Compare with []interface{}.

    We get the account id that we store in the received token and check for the existence of the player.
        If the player exists, continue the process.
        If not, pause and declare "that the player cannot be found."

    Let's take the incoming "room invite id" and check the existence of the room.
        If the room exists, continue the process.
        If there is no room, declare "that the player cannot join a room that does not exist."

    We have the id of the player and with this we will check if the player belongs to a room or not.
        If not, proceed.
        If he/she is, stop and declare "that he/she can only join 1 room at a time."

    Join the room.


**Start play:**
    Firstly check connection is websocket connection.
        If yes, proceed.
        If not, stop and declare "Connection is not supported"

    Before setting up the room, check if the player has the token.
        We get the token from cookies, we check.

 
    We get the account id that we store in the received token and check for the existence of the player.
        If the player exists, continue the process.
        If not, pause and declare "that the player cannot be found".


    We check presence the room?
        If there, proceed.
        If not, pause and declare "Room not found".
        
    We check player owner the room?
        If owner, proceed.
        If not, pause and declare "that player does not owner the this room".

    We check room have a minimum four player?
        If like this, proceed.
        If not, pause and declare "Required min four player for start the game"


    Turn response as cards up to n, and init variables in redis.

    
**Drop card process:**
    We check that the player has played a card in this round.
        If not, proceed.
        If yes, play a card and add it to the discard pile.
    

    We check presence of card.
        If there, proceed.
        If not, pause and declare "this card is undefiend"
    

