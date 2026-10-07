**TCP Communication Message Frames V1**

1. Lock

   1.1. Lock Resource

   1.1.1. Request

        <version>|<messageType>|LOCK|<resourceId>
        1|REQ|LOCK|234324234

   1.1.2. Response

   Successful

         <version>|<messageType>|<status>|LOCK|<resourceId>
         1|RES|SUCCESSFUL|LOCK|234324234

   Failed

         <version>|<messageType>|<status>|LOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|LOCK|234324234|RESOURCE_NOT_FOUND|Resource [234324234] not found.


   1.2. Unlock Resource

   1.2.1. Request

        <version>|<messageType>|UNLOCK|<resourceId>
        1|REQ|UNLOCK|2323232

   1.2.2. Response

   Successful

         <version>|<messageType>|<status>|UNLOCK|<resourceId>
         1|RES|SUCCESS|UNLOCK|2323232

   Failed

         <version>|<messageType>|<status>|UNLOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|UNLOCK|234324234|RESOURCE_NOT_FOUND|Resource [234324234] not found.


2. Stock

   2.1. Create Stock

   2.1.1. Request

         <version>|<messageType>|CREATE_STOCK|<resourceId>|<quantity>
         1|REQ|CREATE_STOCK|12122|100

   2.1.2. Response

   Successful

         <version>|<messageType>|<status>|CREATE_STOCK|<reserveId>
         1|RES|SUCCESS|CREATE_STOCK|5ffe3244-343gf-23232-grg333

   Failed

         <version>|<messageType>|<status>|CREATE_STOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|CREATE_STOCK|232323|STOCK_ALREADY_CREATED|Stock is already created.


   2.2. Reserve Stock

   2.2.1. Request

         <version>|<messageType>|RESERVE_STOCK|<resourceId>|<quantity>
         1|REQ|RESERVE_STOCK|12122|10

   2.2.2. Response

   Successful

         <version>|<messageType>|<status>|RESERVE_STOCK|<reservationId>
         1|RES|SUCCESS|RESERVE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890

   Failed

         <version>|<messageType>|<status>|RESERVE_STOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|RESERVE_STOCK|12122|INSUFFICIENT_STOCK|Insufficient stock for resource [12122].

         1|RES|FAILED|RESERVE_STOCK|12122|STOCK_NOT_FOUND|Stock not found for resource [12122].


   2.3. Release Stock

   2.3.1. Request

         <version>|<messageType>|RELEASE_STOCK|<reservationId>
         1|REQ|RELEASE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890

   2.3.2. Response

   Successful

         <version>|<messageType>|<status>|RELEASE_STOCK|<reservationId>
         1|RES|SUCCESS|RELEASE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890

   Failed

         <version>|<messageType>|<status>|RELEASE_STOCK|<reservationId>|<errorCode>|<errorMessage>
         1|RES|FAILED|RELEASE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890|RESERVATION_NOT_FOUND|Reservation [a1b2c3d4-e5f6-7890-abcd-ef1234567890] not found.

3. Rate Limiter

   3.1. Setup

   3.1.1. Request

         <version>|REQ|RL_SETUP|<resourceId>|<userId>|<timeFrame>|<rate>
         1|REQ|RL_SETUP|resource1|user1|MINUTE|60

   3.1.2. Response

   Successful

         <version>|RES|<status>|RL_SETUP
         1|RES|SUCCESSFUL|RL_SETUP

   Failed

         <version>|RES|FAILED|RL_SETUP|<errorCode>|<errorMessage>

   3.2. Wait 

   3.2.1. Request

         <version>|REQ|RL_WAIT|<resourceId>|userId
         1|REQ|RL_WAIT|resource1|user1

   3.2.2. Response

   Successful

         <version>|RES|SUCCESSFUL|RL_WAIT|<holdFlag>
         1|RES|SUCCESSFUL|RL_WAIT|false

   Failed

         <version>|RES|FAILED|RL_WAIT|<errorCode>|<errorMessage>
         1|RES|RL_WAIT|LIMIT_EXCEEDED|Rate limiter exceeded.
         1|RES|RL_WAIT|TOO_MANY_REQUESTS|Too many wait requests, can not hold any more clients.
