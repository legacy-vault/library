package header

// HTTP Message Header Field Names.

// Permanent.
const (
	HttpHeaderAIM                           = "A-IM"                              // [RFC4229]
	HttpHeaderAccept                        = "Accept"                            // [RFC7231, Section 5.3.2]
	HttpHeaderAcceptAdditions               = "Accept-Additions"                  // [RFC4229]
	HttpHeaderAcceptCharset                 = "Accept-Charset"                    // [RFC7231, Section 5.3.3]
	HttpHeaderAcceptDatetime                = "Accept-Datetime"                   // [RFC7089]
	HttpHeaderAcceptEncoding                = "Accept-Encoding"                   // [RFC7231, Section 5.3.4][RFC7694, Section 3]
	HttpHeaderAcceptFeatures                = "Accept-Features"                   // [RFC4229]
	HttpHeaderAcceptLanguage                = "Accept-Language"                   // [RFC7231, Section 5.3.5]
	HttpHeaderAcceptPatch                   = "Accept-Patch"                      // [RFC5789]
	HttpHeaderAcceptPost                    = "Accept-Post"                       // [https://www.w3.org/TR/ldp/]
	HttpHeaderAcceptRanges                  = "Accept-Ranges"                     // [RFC7233, Section 2.3]
	HttpHeaderAge                           = "Age"                               // [RFC7234, Section 5.1]
	HttpHeaderAllow                         = "Allow"                             // [RFC7231, Section 7.4.1]
	HttpHeaderALPN                          = "ALPN"                              // [RFC7639, Section 2]
	HttpHeaderAltSvc                        = "Alt-Svc"                           // [RFC7838]
	HttpHeaderAltUsed                       = "Alt-Used"                          // [RFC7838]
	HttpHeaderAlternates                    = "Alternates"                        // [RFC4229]
	HttpHeaderApplyToRedirectRef            = "Apply-To-Redirect-Ref"             // [RFC4437]
	HttpHeaderAuthenticationControl         = "Authentication-Control"            // [RFC8053, Section 4]
	HttpHeaderAuthenticationInfo            = "Authentication-Info"               // [RFC7615, Section 3]
	HttpHeaderAuthorization                 = "Authorization"                     // [RFC7235, Section 4.2]
	HttpHeaderCExt                          = "C-Ext"                             // [RFC4229]
	HttpHeaderCMan                          = "C-Man"                             // [RFC4229]
	HttpHeaderCOpt                          = "C-Opt"                             // [RFC4229]
	HttpHeaderCPEP                          = "C-PEP"                             // [RFC4229]
	HttpHeaderCPEPInfo                      = "C-PEP-Info"                        // [RFC4229]
	HttpHeaderCacheControl                  = "Cache-Control"                     // [RFC7234, Section 5.2]
	HttpHeaderCalManagedID                  = "Cal-Managed-ID"                    // [RFC-ietf-calext-caldav-attachments-04, Section 5.1]
	HttpHeaderCalDAVTimezones               = "CalDAV-Timezones"                  // [RFC7809, Section 7.1]
	HttpHeaderCDNLoop                       = "CDN-Loop"                          // [RFC8586]
	HttpHeaderClose                         = "Close"                             // [RFC7230, Section 8.1]
	HttpHeaderConnection                    = "Connection"                        // [RFC7230, Section 6.1]
	HttpHeaderContentBase                   = "Content-Base"                      // [RFC2068][RFC2616]
	HttpHeaderContentDisposition            = "Content-Disposition"               // [RFC6266]
	HttpHeaderContentEncoding               = "Content-Encoding"                  // [RFC7231, Section 3.1.2.2]
	HttpHeaderContentID                     = "Content-ID"                        // [RFC4229]
	HttpHeaderContentLanguage               = "Content-Language"                  // [RFC7231, Section 3.1.3.2]
	HttpHeaderContentLength                 = "Content-Length"                    // [RFC7230, Section 3.3.2]
	HttpHeaderContentLocation               = "Content-Location"                  // [RFC7231, Section 3.1.4.2]
	HttpHeaderContentMD5                    = "Content-MD5"                       // [RFC4229]
	HttpHeaderContentRange                  = "Content-Range"                     // [RFC7233, Section 4.2]
	HttpHeaderContentScriptType             = "Content-Script-Type"               // [RFC4229]
	HttpHeaderContentStyleType              = "Content-Style-Type"                // [RFC4229]
	HttpHeaderContentType                   = "Content-Type"                      // [RFC7231, Section 3.1.1.5]
	HttpHeaderContentVersion                = "Content-Version"                   // [RFC4229]
	HttpHeaderCookie                        = "Cookie"                            // [RFC6265]
	HttpHeaderCookie2                       = "Cookie2"                           // [RFC2965][RFC6265]
	HttpHeaderDASL                          = "DASL"                              // [RFC5323]
	HttpHeaderDAV                           = "DAV"                               // [RFC4918]
	HttpHeaderDate                          = "Date"                              // [RFC7231, Section 7.1.1.2]
	HttpHeaderDefaultStyle                  = "Default-Style"                     // [RFC4229]
	HttpHeaderDeltaBase                     = "Delta-Base"                        // [RFC4229]
	HttpHeaderDepth                         = "Depth"                             // [RFC4918]
	HttpHeaderDerivedFrom                   = "Derived-From"                      // [RFC4229]
	HttpHeaderDestination                   = "Destination"                       // [RFC4918]
	HttpHeaderDifferentialID                = "Differential-ID"                   // [RFC4229]
	HttpHeaderDigest                        = "Digest"                            // [RFC4229]
	HttpHeaderEarlyData                     = "Early-Data"                        // [RFC8470]
	HttpHeaderETag                          = "ETag"                              // [RFC7232, Section 2.3]
	HttpHeaderExpect                        = "Expect"                            // [RFC7231, Section 5.1.1]
	HttpHeaderExpectCT                      = "Expect-CT"                         // [RFC-ietf-httpbis-expect-ct-08]
	HttpHeaderExpires                       = "Expires"                           // [RFC7234, Section 5.3]
	HttpHeaderExt                           = "Ext"                               // [RFC4229]
	HttpHeaderForwarded                     = "Forwarded"                         // [RFC7239]
	HttpHeaderFrom                          = "From"                              // [RFC7231, Section 5.5.1]
	HttpHeaderGetProfile                    = "GetProfile"                        // [RFC4229]
	HttpHeaderHobareg                       = "Hobareg"                           // [RFC7486, Section 6.1.1]
	HttpHeaderHost                          = "Host"                              // [RFC7230, Section 5.4]
	HttpHeaderHTTP2Settings                 = "HTTP2-Settings"                    // [RFC7540, Section 3.2.1]
	HttpHeaderIM                            = "IM"                                // [RFC4229]
	HttpHeaderIf                            = "If"                                // [RFC4918]
	HttpHeaderIfMatch                       = "If-Match"                          // [RFC7232, Section 3.1]
	HttpHeaderIfModifiedSince               = "If-Modified-Since"                 // [RFC7232, Section 3.3]
	HttpHeaderIfNoneMatch                   = "If-None-Match"                     // [RFC7232, Section 3.2]
	HttpHeaderIfRange                       = "If-Range"                          // [RFC7233, Section 3.2]
	HttpHeaderIfScheduleTagMatch            = "If-Schedule-Tag-Match"             // [RFC6638]
	HttpHeaderIfUnmodifiedSince             = "If-Unmodified-Since"               // [RFC7232, Section 3.4]
	HttpHeaderIncludeReferredTokenBindingID = "Include-Referred-Token-Binding-ID" // [RFC8473]
	HttpHeaderKeepAlive                     = "Keep-Alive"                        // [RFC4229]
	HttpHeaderLabel                         = "Label"                             // [RFC4229]
	HttpHeaderLastModified                  = "Last-Modified"                     // [RFC7232, Section 2.2]
	HttpHeaderLink                          = "Link"                              // [RFC8288]
	HttpHeaderLocation                      = "Location"                          // [RFC7231, Section 7.1.2]
	HttpHeaderLockToken                     = "Lock-Token"                        // [RFC4918]
	HttpHeaderMan                           = "Man"                               // [RFC4229]
	HttpHeaderMaxForwards                   = "Max-Forwards"                      // [RFC7231, Section 5.1.2]
	HttpHeaderMementoDatetime               = "Memento-Datetime"                  // [RFC7089]
	HttpHeaderMeter                         = "Meter"                             // [RFC4229]
	HttpHeaderMIMEVersion                   = "MIME-Version"                      // [RFC7231, Appendix A.1]
	HttpHeaderNegotiate                     = "Negotiate"                         // [RFC4229]
	HttpHeaderOpt                           = "Opt"                               // [RFC4229]
	HttpHeaderOptionalWWWAuthenticate       = "Optional-WWW-Authenticate"         // [RFC8053, Section 3]
	HttpHeaderOrderingType                  = "Ordering-Type"                     // [RFC4229]
	HttpHeaderOrigin                        = "Origin"                            // [RFC6454]
	HttpHeaderOSCORE                        = "OSCORE"                            // [RFC-ietf-core-object-security-16, Section 11.1]
	HttpHeaderOverwrite                     = "Overwrite"                         // [RFC4918]
	HttpHeaderP3P                           = "P3P"                               // [RFC4229]
	HttpHeaderPEP                           = "PEP"                               // [RFC4229]
	HttpHeaderPICSLabel                     = "PICS-Label"                        // [RFC4229]
	HttpHeaderPepInfo                       = "Pep-Info"                          // [RFC4229]
	HttpHeaderPosition                      = "Position"                          // [RFC4229]
	HttpHeaderPragma                        = "Pragma"                            // [RFC7234, Section 5.4]
	HttpHeaderPrefer                        = "Prefer"                            // [RFC7240]
	HttpHeaderPreferenceApplied             = "Preference-Applied"                // [RFC7240]
	HttpHeaderProfileObject                 = "ProfileObject"                     // [RFC4229]
	HttpHeaderProtocol                      = "Protocol"                          // [RFC4229]
	HttpHeaderProtocolInfo                  = "Protocol-Info"                     // [RFC4229]
	HttpHeaderProtocolQuery                 = "Protocol-Query"                    // [RFC4229]
	HttpHeaderProtocolRequest               = "Protocol-Request"                  // [RFC4229]
	HttpHeaderProxyAuthenticate             = "Proxy-Authenticate"                // [RFC7235, Section 4.3]
	HttpHeaderProxyAuthenticationInfo       = "Proxy-Authentication-Info"         // [RFC7615, Section 4]
	HttpHeaderProxyAuthorization            = "Proxy-Authorization"               // [RFC7235, Section 4.4]
	HttpHeaderProxyFeatures                 = "Proxy-Features"                    // [RFC4229]
	HttpHeaderProxyInstruction              = "Proxy-Instruction"                 // [RFC4229]
	HttpHeaderPublic                        = "Public"                            // [RFC4229]
	HttpHeaderPublicKeyPins                 = "Public-Key-Pins"                   // [RFC7469]
	HttpHeaderPublicKeyPinsReportOnly       = "Public-Key-Pins-Report-Only"       // [RFC7469]
	HttpHeaderRange                         = "Range"                             // [RFC7233, Section 3.1]
	HttpHeaderRedirectRef                   = "Redirect-Ref"                      // [RFC4437]
	HttpHeaderReferer                       = "Referer"                           // [RFC7231, Section 5.5.2]
	HttpHeaderReplayNonce                   = "Replay-Nonce"                      // [RFC8555, Section 6.5.1]
	HttpHeaderRetryAfter                    = "Retry-After"                       // [RFC7231, Section 7.1.3]
	HttpHeaderSafe                          = "Safe"                              // [RFC4229]
	HttpHeaderScheduleReply                 = "Schedule-Reply"                    // [RFC6638]
	HttpHeaderScheduleTag                   = "Schedule-Tag"                      // [RFC6638]
	HttpHeaderSecTokenBinding               = "Sec-Token-Binding"                 // [RFC8473]
	HttpHeaderSecWebSocketAccept            = "Sec-WebSocket-Accept"              // [RFC6455]
	HttpHeaderSecWebSocketExtensions        = "Sec-WebSocket-Extensions"          // [RFC6455]
	HttpHeaderSecWebSocketKey               = "Sec-WebSocket-Key"                 // [RFC6455]
	HttpHeaderSecWebSocketProtocol          = "Sec-WebSocket-Protocol"            // [RFC6455]
	HttpHeaderSecWebSocketVersion           = "Sec-WebSocket-Version"             // [RFC6455]
	HttpHeaderSecurityScheme                = "Security-Scheme"                   // [RFC4229]
	HttpHeaderServer                        = "Server"                            // [RFC7231, Section 7.4.2]
	HttpHeaderSetCookie                     = "Set-Cookie"                        // [RFC6265]
	HttpHeaderSetCookie2                    = "Set-Cookie2"                       // [RFC2965][RFC6265]
	HttpHeaderSetProfile                    = "SetProfile"                        // [RFC4229]
	HttpHeaderSLUG                          = "SLUG"                              // [RFC5023]
	HttpHeaderSoapAction                    = "SoapAction"                        // [RFC4229]
	HttpHeaderStatusURI                     = "Status-URI"                        // [RFC4229]
	HttpHeaderStrictTransportSecurity       = "Strict-Transport-Security"         // [RFC6797]
	HttpHeaderSunset                        = "Sunset"                            // [RFC8594]
	HttpHeaderSurrogateCapability           = "Surrogate-Capability"              // [RFC4229]
	HttpHeaderSurrogateControl              = "Surrogate-Control"                 // [RFC4229]
	HttpHeaderTCN                           = "TCN"                               // [RFC4229]
	HttpHeaderTE                            = "TE"                                // [RFC7230, Section 4.3]
	HttpHeaderTimeout                       = "Timeout"                           // [RFC4918]
	HttpHeaderTopic                         = "Topic"                             // [RFC8030, Section 5.4]
	HttpHeaderTrailer                       = "Trailer"                           // [RFC7230, Section 4.4]
	HttpHeaderTransferEncoding              = "Transfer-Encoding"                 // [RFC7230, Section 3.3.1]
	HttpHeaderTTL                           = "TTL"                               // [RFC8030, Section 5.2]
	HttpHeaderUrgency                       = "Urgency"                           // [RFC8030, Section 5.3]
	HttpHeaderURI                           = "URI"                               // [RFC4229]
	HttpHeaderUpgrade                       = "Upgrade"                           // [RFC7230, Section 6.7]
	HttpHeaderUserAgent                     = "User-Agent"                        // [RFC7231, Section 5.5.3]
	HttpHeaderVariantVary                   = "Variant-Vary"                      // [RFC4229]
	HttpHeaderVary                          = "Vary"                              // [RFC7231, Section 7.1.4]
	HttpHeaderVia                           = "Via"                               // [RFC7230, Section 5.7.1]
	HttpHeaderWWWAuthenticate               = "WWW-Authenticate"                  // [RFC7235, Section 4.1]
	HttpHeaderWantDigest                    = "Want-Digest"                       // [RFC4229]
	HttpHeaderWarning                       = "Warning"                           // [RFC7234, Section 5.5]
	HttpHeaderXContentTypeOptions           = "X-Content-Type-Options"            // [https://fetch.spec.whatwg.org/#x-content-type-options-header]
	HttpHeaderXFrameOptions                 = "X-Frame-Options"                   // [RFC7034]
)

// Provisional.
const (
	HttpHeaderAccessControl                 = "Access-Control"                   // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials" // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"     // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"     // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"      // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlMaxAge           = "Access-Control-Max-Age"           // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlRequestMethod    = "Access-Control-Request-Method"    // [W3C Web Application Formats Working Group]
	HttpHeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"   // [W3C Web Application Formats Working Group]
	HttpHeaderAMPCacheTransform             = "AMP-Cache-Transform"              // [https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-transform.md]
	HttpHeaderCompliance                    = "Compliance"                       // [RFC4229]
	HttpHeaderContentTransferEncoding       = "Content-Transfer-Encoding"        // [RFC4229]
	HttpHeaderCost                          = "Cost"                             // [RFC4229]
	HttpHeaderEDIINTFeatures                = "EDIINT-Features"                  // [RFC6017]
	HttpHeaderMessageID                     = "Message-ID"                       // [RFC4229]
	HttpHeaderMethodCheck                   = "Method-Check"                     // [W3C Web Application Formats Working Group]
	HttpHeaderMethodCheckExpires            = "Method-Check-Expires"             // [W3C Web Application Formats Working Group]
	HttpHeaderNonCompliance                 = "Non-Compliance"                   // [RFC4229]
	HttpHeaderOptional                      = "Optional"                         // [RFC4229]
	HttpHeaderRefererRoot                   = "Referer-Root"                     // [W3C Web Application Formats Working Group]
	HttpHeaderResolutionHint                = "Resolution-Hint"                  // [RFC4229]
	HttpHeaderResolverLocation              = "Resolver-Location"                // [RFC4229]
	HttpHeaderSubOK                         = "SubOK"                            // [RFC4229]
	HttpHeaderSubst                         = "Subst"                            // [RFC4229]
	HttpHeaderTimingAllowOrigin             = "Timing-Allow-Origin"              // [https://www.w3.org/TR/resource-timing-1/#timing-allow-origin]
	HttpHeaderTitle                         = "Title"                            // [RFC4229]
	HttpHeaderTraceparent                   = "Traceparent"                      // [https://www.w3.org/TR/trace-context/#traceparent-field]
	HttpHeaderTracestate                    = "Tracestate"                       // [https://www.w3.org/TR/trace-context/#tracestate-field]
	HttpHeaderUAColor                       = "UA-Color"                         // [RFC4229]
	HttpHeaderUAMedia                       = "UA-Media"                         // [RFC4229]
	HttpHeaderUAPixels                      = "UA-Pixels"                        // [RFC4229]
	HttpHeaderUAResolution                  = "UA-Resolution"                    // [RFC4229]
	HttpHeaderUAWindowpixels                = "UA-Windowpixels"                  // [RFC4229]
	HttpHeaderVersion                       = "Version"                          // [RFC4229]
	HttpHeaderXDeviceAccept                 = "X-Device-Accept"                  // [W3C Mobile Web Best Practices Working Group]
	HttpHeaderXDeviceAcceptCharset          = "X-Device-Accept-Charset"          // [W3C Mobile Web Best Practices Working Group]
	HttpHeaderXDeviceAcceptEncoding         = "X-Device-Accept-Encoding"         // [W3C Mobile Web Best Practices Working Group]
	HttpHeaderXDeviceAcceptLanguage         = "X-Device-Accept-Language"         // [W3C Mobile Web Best Practices Working Group]
	HttpHeaderXDeviceUserAgent              = "X-Device-User-Agent"              // [W3C Mobile Web Best Practices Working Group]
)
