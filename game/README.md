**Endpoints:**
    root endpoint: game -> /game/?

    POST -> game/room/create-room -> For create a room
        requirements body: nil

    POST -> game/room/join-room -> For join a room
        requirements body: RoomId

    GET ->  game/items/rooms -> For get all rooms
        requirements token: Access with access key


**Create room:**
    Before creating the room, we check if the user has the token.
        We get the token from cookies, we check.

    Based on the parameters expected from the client, the parameters of the previously created data structure are compared with the incoming parameters.
        Get the HTTP request body.
        Compare with []interface{}.

    We get the account id that we store in the received token and check for the existence of the user.
        If the user exists, continue the process.
        If not, pause and declare that the user cannot be found.

    We have the id of the user and with this we will check if the user belongs to a room or not.
        If not, proceed.
        If he/she is, pause and declare that he/she can only join 1 room at a time.

    Create the room and then join the user to the room.

**Join room:**
    Before setting up the room, check if the user has the token.
        We get the token from cookies, we check.

    Based on the parameters expected from the client, the parameters of the previously created data structure are compared with the incoming parameters.
        Get the HTTP request body.
        Compare with []interface{}.

    We get the account id that we store in the received token and check for the existence of the user.
        If the user exists, continue the process.
        If not, pause and declare that the user cannot be found.

    Let's take the incoming room id and check the existence of the room.
        If the room exists, continue the process.
        If there is no room, declare that the user cannot join a room that does not exist.

    We have the id of the user and with this we will check if the user belongs to a room or not.
        If not, proceed.
        If he/she is, stop and declare that he/she can only join 1 room at a time.

    Join the room.