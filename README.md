

**TCP Communication Message Frames V1**

1. Lock Resource

   1.1. Request

        <version>|<messageType>|LOCK|<resourceId>
        1|REQ|LOCK|234324234

   1.2. Response
   
   Successful

         <version>|<messageType>|<status>|LOCK|<resourceId>
         <1>|RES|SUCCESSFUL|LOCK|234324234

   Failed

         <version>|<messageType>|<status>|LOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|LOCK|234324234|RESOURCE_NOT_FOUND|Resource [234324234] not found.

         
2. Unlock Resource

   2.1. Request

        <version>|<messageType>|UNLOCK|<resourceId>
        1|REQ|UNLOCK|2323232

   2.2. Response

   Successful

         <version>|<messageType>|<status>|UNLOCK|<resourceId>
         1|RES|SUCCESS|2323232

   Failed

         <version>|<messageType>|<status>|UNLOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|UNLOCK|234324234|RESOURCE_NOT_FOUND|Resource [234324234] not found.


3. Create Stock

   3.1. Request

         <version>|<messageType>|CREATE_STOCK|<resourceId>|<quantity>
         1|REQ|CREATE_STOCK|12122|100

   3.2. Response

   Successful

         <version>|<messageType>|<status>|CREATE_STOCK|<reserveId>
         <1>|RES|SUCCESS|CREATE_STOCK|5ffe3244-343gf-23232-grg333

   Failed

         <version>|<messageType>|<status>|CREATE_STOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|CREATE_STOCK|232323|STOCK_ALREADY_CREATED|Stock is already created.


4. Reserve Stock

   4.1. Request

         <version>|<messageType>|RESERVE_STOCK|<resourceId>|<quantity>
         1|REQ|RESERVE_STOCK|12122|10

   4.2. Response

   Successful

         <version>|<messageType>|<status>|RESERVE_STOCK|<reservationId>
         1|RES|SUCCESS|RESERVE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890

   Failed

         <version>|<messageType>|<status>|RESERVE_STOCK|<resourceId>|<errorCode>|<errorMessage>
         1|RES|FAILED|RESERVE_STOCK|12122|INSUFFICIENT_STOCK|Insufficient stock for resource [12122].

         1|RES|FAILED|RESERVE_STOCK|12122|STOCK_NOT_FOUND|Stock not found for resource [12122].


5. Release Stock

   5.1. Request

         <version>|<messageType>|RELEASE_STOCK|<reservationId>
         1|REQ|RELEASE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890

   5.2. Response

   Successful

         <version>|<messageType>|<status>|RELEASE_STOCK|<reservationId>
         1|RES|SUCCESS|RELEASE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890

   Failed

         <version>|<messageType>|<status>|RELEASE_STOCK|<reservationId>|<errorCode>|<errorMessage>
         1|RES|FAILED|RELEASE_STOCK|a1b2c3d4-e5f6-7890-abcd-ef1234567890|RESERVATION_NOT_FOUND|Reservation [a1b2c3d4-e5f6-7890-abcd-ef1234567890] not found.



