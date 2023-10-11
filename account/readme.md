Account interesting with user informations like username, password, which support login, register for now.

Endpoints:

Root endpoint: account -> /account/?

        POST -> account/register -> For create an account
            requirements body: keywords(Username, Password) = "test123", "pass123321".

        POST -> account/login -> For login to system/game
            requirements body: keywords(Username, Password) = "test123", "pass123321".

        GET ->  account/logout -> For leave the system
            requirements token: You are must be logged in.

        GET ->  account/items/room -> For get room of account
            requirements token: You are must be logged in.
