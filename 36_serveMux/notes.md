# Clarifications

Entry point -> net/http package

<ol>

<li> There is a type "Handler" an interface
that has an ServeHttp method that has ResponseWriter, *Request as its parameters.

<ul>
    <li>a request is a struct with different fields</li>
    <li>ResponseWriter allows us to write some headers in the response</li>
</ul>
<li>
ListenAndServe takes a "Handler" interface
</li>

<li>For routing we have to consider: NewServeMux returns a pointer to serveMux -> *ServeMux

Any value of type serveMux implements the Handler interface -> which means when we create a ListenAndServe we can
pass the *ServeMux type.
*ServeMux has several other methods attached to it:

* Handle -> takes a Handler
* HandleFunc -> takes a function with the parameters ResponseWriter and *Request. That is not a type handler but simply a function
* Handler
* ServeHTTP
</li>
<li>
We can also pass "nil" to the ListenAndServe method, in that case we can attach routes by using:
    * http.Handle -> receives a pattern and a type that implements the Handler interface (which means that that type implement the ServeHTTP method)
    * http.HandleFunc -> receives a patterns and a signature -> signature is -> func(res http.ResponseWriter, req *http.Request)
</li>
</ol>

<strong>Note:</strong><br>
There also exist the type "HandlerFunc" whose UNDERLYING type is func(ResponseWriter, *Request). That type also implements the class ServeHTTP thus it is also of type "Handler"
