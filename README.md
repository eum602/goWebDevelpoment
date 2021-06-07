# Sessions

Unique id is often called session id.
For any given session there is which user (through a userId) that has that session. With that userId all other information can be obtained from a storage like a database. There are two separate different storage areas, one for the sessions (sessionId -> userId) and the other area where there is a map between a userId and the user Data (userId -> userData)

