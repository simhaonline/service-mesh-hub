
---
title: "routing.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `config.zephyr.solo.io` 
#### Types:


- [RoutingRuleSpec](#routingrulespec)
- [RoutingRuleStatus](#routingrulestatus)
- [RetryPolicy](#retrypolicy)
- [MultiDestination](#multidestination)
- [WeightedDestination](#weighteddestination)
- [FaultInjection](#faultinjection)
- [Delay](#delay)
- [Abort](#abort)
- [HeaderManipulation](#headermanipulation)
- [CorsPolicy](#corspolicy)
- [Matcher](#matcher)
- [StringMatch](#stringmatch)
- [HeaderMatcher](#headermatcher)
- [QueryParameterMatcher](#queryparametermatcher)
  



##### Source File: [github.com/solo-io/mesh-projects/api/config/v1alpha1/routing.proto](https://github.com/solo-io/mesh-projects/blob/master/api/config/v1alpha1/routing.proto)





---
### RoutingRuleSpec

 
a routing rule applies some L7 routing features to an existing mesh
routing rules specify the following:
for all requests:
- originating from from **source pods**
- sent to **destination pods**
- matching one or more **request matcher**
apply the specified RoutingRuleSpec
the routing configuration that will be applied to the mesh(es)

```yaml
"sourceSelector": .core.zephyr.solo.io.Selector
"destinationSelector": .core.zephyr.solo.io.Selector
"requestMatchers": .config.zephyr.solo.io.Matcher
"trafficShift": .config.zephyr.solo.io.MultiDestination
"faultInjection": .config.zephyr.solo.io.FaultInjection
"requestTimeout": .google.protobuf.Duration
"retries": .config.zephyr.solo.io.RetryPolicy
"corsPolicy": .config.zephyr.solo.io.CorsPolicy
"mirror": .core.zephyr.solo.io.ResourceRef
"headerManipulation": .config.zephyr.solo.io.HeaderManipulation

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `sourceSelector` | [.core.zephyr.solo.io.Selector](../../../core/v1alpha1/selector.proto.sk/#selector) | requests originating from these pods will have the rule applied leave empty to have all pods in the mesh apply these rules > Note: Source Selectors are ignored when RoutingRules are applied to pods in a Linkerd mesh. RoutingRules will apply to all selected destinations in Linkerd, regardless of the source. |  |
| `destinationSelector` | [.core.zephyr.solo.io.Selector](../../../core/v1alpha1/selector.proto.sk/#selector) | requests destined for these pods will have the rule applied leave empty to apply to all destination pods in the mesh. |  |
| `requestMatchers` | [.config.zephyr.solo.io.Matcher](../routing.proto.sk/#matcher) | if specified, this rule will only apply to http requests in the mesh matching these parameters note that Linkerd only supports matching on Request Path and Method. |  |
| `trafficShift` | [.config.zephyr.solo.io.MultiDestination](../routing.proto.sk/#multidestination) | enables traffic shifting, i.e. to reroute requests to a different service, to a subset of pods based on their label, and/or split traffic between multiple services. Only one of `trafficShift`, `faultInjection`, `requestTimeout`, `retries`, `corsPolicy`, or `headerManipulation` can be set. |  |
| `faultInjection` | [.config.zephyr.solo.io.FaultInjection](../routing.proto.sk/#faultinjection) | enable fault injection on requests. Only one of `faultInjection`, `trafficShift`, `requestTimeout`, `retries`, `corsPolicy`, or `headerManipulation` can be set. |  |
| `requestTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | set a timeout on requests. Only one of `requestTimeout`, `trafficShift`, `faultInjection`, `retries`, `corsPolicy`, or `headerManipulation` can be set. |  |
| `retries` | [.config.zephyr.solo.io.RetryPolicy](../routing.proto.sk/#retrypolicy) | set a retry policy on requests. Only one of `retries`, `trafficShift`, `faultInjection`, `requestTimeout`, `corsPolicy`, or `headerManipulation` can be set. |  |
| `corsPolicy` | [.config.zephyr.solo.io.CorsPolicy](../routing.proto.sk/#corspolicy) | set a Cross-Origin Resource Sharing policy (CORS) for requests. Refer to https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS for further details about cross origin resource sharing. Only one of `corsPolicy`, `trafficShift`, `faultInjection`, `requestTimeout`, `retries`, or `headerManipulation` can be set. |  |
| `mirror` | [.core.zephyr.solo.io.ResourceRef](../../../core/v1alpha1/ref.proto.sk/#resourceref) | Mirror HTTP traffic to a another destination. Traffic will still be sent to its original destination as normal. Only one of `mirror`, `trafficShift`, `faultInjection`, `requestTimeout`, `retries`, or `headerManipulation` can be set. |  |
| `headerManipulation` | [.config.zephyr.solo.io.HeaderManipulation](../routing.proto.sk/#headermanipulation) | manipulate request and response headers. Only one of `headerManipulation`, `trafficShift`, `faultInjection`, `requestTimeout`, `retries`, or `mirror` can be set. |  |




---
### RoutingRuleStatus



```yaml

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 




---
### RetryPolicy

 
RetryPolicy contains mesh-specific retry configuration
Different meshes support different Retry features
Service Mesh Hub's RetryPolicy exposes config for multiple meshes simultaneously,
Allowing the same RoutingRule to apply retries to different mesh types
The configuration applied to the target mesh will use the corresponding
config for each type, while other config types will be ignored

```yaml
"attempts": int
"perTryTimeout": .google.protobuf.Duration

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `attempts` | `int` | Number of retries for a given request. |  |
| `perTryTimeout` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | Timeout per retry attempt for a given request. format: 1h/1m/1s/1ms. MUST BE >=1ms. |  |




---
### MultiDestination



```yaml
"destinations": []config.zephyr.solo.io.MultiDestination.WeightedDestination

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `destinations` | [[]config.zephyr.solo.io.MultiDestination.WeightedDestination](../routing.proto.sk/#weighteddestination) |  |  |




---
### WeightedDestination



```yaml
"destination": .core.zephyr.solo.io.ResourceRef
"weight": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `destination` | [.core.zephyr.solo.io.ResourceRef](../../../core/v1alpha1/ref.proto.sk/#resourceref) |  |  |
| `weight` | `int` | Routing to each destination will be balanced by the ratio of the destination's weight to the total weight on a route. |  |




---
### FaultInjection

 
FaultInjection can be used to specify one or more faults to inject
while forwarding http requests to the destination specified in a route.
Faults include aborting the Http request from downstream service, and/or delaying
proxying of requests. A fault rule MUST HAVE delay or abort.

```yaml
"delay": .config.zephyr.solo.io.FaultInjection.Delay
"abort": .config.zephyr.solo.io.FaultInjection.Abort
"percentage": float

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `delay` | [.config.zephyr.solo.io.FaultInjection.Delay](../routing.proto.sk/#delay) | Delay requests before forwarding, emulating various failures such as network issues, overloaded upstream service, etc. Only one of `delay` or `abort` can be set. |  |
| `abort` | [.config.zephyr.solo.io.FaultInjection.Abort](../routing.proto.sk/#abort) | Abort Http request attempts and return error codes back to downstream service, giving the impression that the upstream service is faulty. Only one of `abort` or `delay` can be set. |  |
| `percentage` | `float` | Percentage of requests to be faulted with the error code provided. Values range between 0 and 100. |  |




---
### Delay

 
The _fixedDelay_ field is used to indicate the amount of delay in seconds.
The optional _percentage_ field can be used to only delay a certain
percentage of requests. If left unspecified, all request will be delayed.

```yaml
"fixedDelay": .google.protobuf.Duration
"exponentialDelay": .google.protobuf.Duration
"percentage": float

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `fixedDelay` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | Add a fixed delay before forwarding the request. Format: 1h/1m/1s/1ms. MUST be >=1ms. Only one of `fixedDelay` or `exponentialDelay` can be set. |  |
| `exponentialDelay` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | $hide_from_docs. Only one of `exponentialDelay` or `fixedDelay` can be set. |  |
| `percentage` | `float` | Percentage of requests on which the delay will be injected. value from 0.0 to 100.0. |  |




---
### Abort

 
The _httpStatus_ field is used to indicate the HTTP status code to
return to the caller. The optional _percentage_ field can be used to only
abort a certain percentage of requests. If not specified, all requests are
aborted.

```yaml
"httpStatus": int

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `httpStatus` | `int` | REQUIRED. HTTP status code to use to abort the Http request. |  |




---
### HeaderManipulation

 
manipulate request and response headers

```yaml
"removeResponseHeaders": []string
"appendResponseHeaders": map<string, string>
"removeRequestHeaders": []string
"appendRequestHeaders": map<string, string>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `removeResponseHeaders` | `[]string` | HTTP headers to remove before returning a response to the caller. |  |
| `appendResponseHeaders` | `map<string, string>` | Additional HTTP headers to add before returning a response to the caller. |  |
| `removeRequestHeaders` | `[]string` | HTTP headers to remove before forwarding a request to the destination service. |  |
| `appendRequestHeaders` | `map<string, string>` | Additional HTTP headers to add before forwarding a request to the destination service. |  |




---
### CorsPolicy



```yaml
"allowOrigin": []string
"allowOrigins": []config.zephyr.solo.io.StringMatch
"allowMethods": []string
"allowHeaders": []string
"exposeHeaders": []string
"maxAge": .google.protobuf.Duration
"allowCredentials": .google.protobuf.BoolValue

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `allowOrigin` | `[]string` | The list of origins that are allowed to perform CORS requests. The content will be serialized into the Access-Control-Allow-Origin header. Wildcard * will allow all origins. $hide_from_docs. |  |
| `allowOrigins` | [[]config.zephyr.solo.io.StringMatch](../routing.proto.sk/#stringmatch) | String patterns that match allowed origins. An origin is allowed if any of the string matchers match. If a match is found, then the outgoing Access-Control-Allow-Origin would be set to the origin as provided by the client. |  |
| `allowMethods` | `[]string` | List of HTTP methods allowed to access the resource. The content will be serialized into the Access-Control-Allow-Methods header. |  |
| `allowHeaders` | `[]string` | List of HTTP headers that can be used when requesting the resource. Serialized to Access-Control-Allow-Headers header. |  |
| `exposeHeaders` | `[]string` | A white list of HTTP headers that the browsers are allowed to access. Serialized into Access-Control-Expose-Headers header. |  |
| `maxAge` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | Specifies how long the results of a preflight request can be cached. Translates to the `Access-Control-Max-Age` header. |  |
| `allowCredentials` | [.google.protobuf.BoolValue](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/bool-value) | Indicates whether the caller is allowed to send the actual request (not the preflight) using credentials. Translates to `Access-Control-Allow-Credentials` header. |  |




---
### Matcher

 
Parameters for matching routes

```yaml
"prefix": string
"exact": string
"regex": string
"headers": []config.zephyr.solo.io.HeaderMatcher
"queryParameters": []config.zephyr.solo.io.QueryParameterMatcher
"methods": []string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `prefix` | `string` | If specified, the route is a prefix rule meaning that the prefix must match the beginning of the *:path* header. Only one of `prefix`, or `regex` can be set. |  |
| `exact` | `string` | If specified, the route is an exact path rule meaning that the path must exactly match the *:path* header once the query string is removed. Only one of `exact`, or `regex` can be set. |  |
| `regex` | `string` | If specified, the route is a regular expression rule meaning that the regex must match the *:path* header once the query string is removed. The entire path (without the query string) must match the regex. The rule will not match if only a sub-sequence of the *:path* header matches the regex. The regex grammar is defined `here <http://en.cppreference.com/w/cpp/regex/ecmascript>`_. Only one of `regex`, or `exact` can be set. |  |
| `headers` | [[]config.zephyr.solo.io.HeaderMatcher](../routing.proto.sk/#headermatcher) | Specifies a set of headers that the route should match on. The router will check the request’s headers against all the specified headers in the route config. A match will happen if all the headers in the route are present in the request with the same values (or based on presence if the value field is not in the config). |  |
| `queryParameters` | [[]config.zephyr.solo.io.QueryParameterMatcher](../routing.proto.sk/#queryparametermatcher) | Specifies a set of URL query parameters on which the route should match. The router will check the query string from the *path* header against all the specified query parameters. If the number of specified query parameters is nonzero, they all must match the *path* header's query string for a match to occur. |  |
| `methods` | `[]string` | HTTP Method/Verb(s) to match on. If none specified, the matcher will ignore the HTTP Method. |  |




---
### StringMatch

 
Describes how to match a given string in HTTP headers. Match is
case-sensitive.

```yaml
"exact": string
"prefix": string
"regex": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `exact` | `string` | exact string match. Only one of `exact`, or `regex` can be set. |  |
| `prefix` | `string` | prefix-based match. Only one of `prefix`, or `regex` can be set. |  |
| `regex` | `string` | ECMAscript style regex-based match. Only one of `regex`, or `prefix` can be set. |  |




---
### HeaderMatcher



```yaml
"name": string
"value": string
"regex": bool
"invertMatch": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` | Specifies the name of the header in the request. Only one of `name`, or `regex` can be set. |  |
| `value` | `string` | Specifies the value of the header. If the value is absent a request that has the name header will match, regardless of the header’s value. Only one of `value`, or `regex` can be set. |  |
| `regex` | `bool` | Specifies whether the header value should be treated as regex or not. Only one of `regex`, or `value` can be set. |  |
| `invertMatch` | `bool` | If set to true, the result of the match will be inverted. Defaults to false. Examples: * name=foo, invert_match=true: matches if no header named `foo` is present * name=foo, value=bar, invert_match=true: matches if no header named `foo` with value `bar` is present * name=foo, value=``\d{3}``, regex=true, invert_match=true: matches if no header named `foo` with a value consisting of three integers is present. |  |




---
### QueryParameterMatcher

 
Query parameter matching treats the query string of a request's :path header
as an ampersand-separated list of keys and/or key=value elements.

```yaml
"name": string
"value": string
"regex": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` | Specifies the name of a key that must be present in the requested *path*'s query string. Only one of `name`, or `regex` can be set. |  |
| `value` | `string` | Specifies the value of the key. If the value is absent, a request that contains the key in its query string will match, whether the key appears with a value (e.g., "?debug=true") or not (e.g., "?debug"). Only one of `value`, or `regex` can be set. |  |
| `regex` | `bool` | Specifies whether the query parameter value is a regular expression. Defaults to false. The entire query parameter value (i.e., the part to the right of the equals sign in "key=value") must match the regex. E.g., the regex "\d+$" will match "123" but not "a123" or "123a". Only one of `regex`, or `value` can be set. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->