

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
