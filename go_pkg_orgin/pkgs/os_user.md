

# user
`import "os/user"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a id="pkg-overview">Overview</a>
Package user allows user account lookups by name or id.

For most Unix systems, this package has two internal implementations of
resolving user and group ids to names. One is written in pure Go and
parses /etc/passwd and /etc/group. The other is cgo-based and relies on
the standard C library (libc) routines such as getpwuid_r and getgrnam_r.

When cgo is available, cgo-based (libc-backed) code is used by default.
This can be overridden by using osusergo build tag, which enforces
the pure Go implementation.




## <a id="pkg-index">Index</a>
* [type Group](#Group)
  * [func LookupGroup(name string) (*Group, error)](#LookupGroup)
  * [func LookupGroupId(gid string) (*Group, error)](#LookupGroupId)
* [type UnknownGroupError](#UnknownGroupError)
  * [func (e UnknownGroupError) Error() string](#UnknownGroupError.Error)
* [type UnknownGroupIdError](#UnknownGroupIdError)
  * [func (e UnknownGroupIdError) Error() string](#UnknownGroupIdError.Error)
* [type UnknownUserError](#UnknownUserError)
  * [func (e UnknownUserError) Error() string](#UnknownUserError.Error)
* [type UnknownUserIdError](#UnknownUserIdError)
  * [func (e UnknownUserIdError) Error() string](#UnknownUserIdError.Error)
* [type User](#User)
  * [func Current() (*User, error)](#Current)
  * [func Lookup(username string) (*User, error)](#Lookup)
  * [func LookupId(uid string) (*User, error)](#LookupId)
  * [func (u *User) GroupIds() ([]string, error)](#User.GroupIds)




#### <a id="pkg-files">Package files</a>
[cgo_lookup_unix.go](https://golang.org/src/os/user/cgo_lookup_unix.go) [getgrouplist_unix.go](https://golang.org/src/os/user/getgrouplist_unix.go) [listgroups_unix.go](https://golang.org/src/os/user/listgroups_unix.go) [lookup.go](https://golang.org/src/os/user/lookup.go) [user.go](https://golang.org/src/os/user/user.go) 








## <a id="Group">type</a> [Group](https://golang.org/src/os/user/user.go?s=1880:1953#L46)
Group represents a grouping of users.

On POSIX systems Gid contains a decimal number representing the group ID.


<pre>type Group struct {
<span id="Group.Gid"></span>    Gid  <a href="/pkg/builtin/#string">string</a> <span class="comment">// group ID</span>
<span id="Group.Name"></span>    Name <a href="/pkg/builtin/#string">string</a> <span class="comment">// group name</span>
}
</pre>









### <a id="LookupGroup">func</a> [LookupGroup](https://golang.org/src/os/user/lookup.go?s=1320:1365#L40)
<pre>func LookupGroup(name <a href="/pkg/builtin/#string">string</a>) (*<a href="#Group">Group</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupGroup looks up a group by name. If the group cannot be found, the
returned error is of type UnknownGroupError.




### <a id="LookupGroupId">func</a> [LookupGroupId](https://golang.org/src/os/user/lookup.go?s=1527:1573#L46)
<pre>func LookupGroupId(gid <a href="/pkg/builtin/#string">string</a>) (*<a href="#Group">Group</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupGroupId looks up a group by groupid. If the group cannot be found, the
returned error is of type UnknownGroupIdError.






## <a id="UnknownGroupError">type</a> [UnknownGroupError](https://golang.org/src/os/user/user.go?s=2648:2677#L76)
UnknownGroupError is returned by LookupGroup when
a group cannot be found.


<pre>type UnknownGroupError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnknownGroupError.Error">func</a> (UnknownGroupError) [Error](https://golang.org/src/os/user/user.go?s=2679:2720#L78)
<pre>func (e <a href="#UnknownGroupError">UnknownGroupError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnknownGroupIdError">type</a> [UnknownGroupIdError](https://golang.org/src/os/user/user.go?s=2439:2470#L68)
UnknownGroupIdError is returned by LookupGroupId when
a group cannot be found.


<pre>type UnknownGroupIdError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnknownGroupIdError.Error">func</a> (UnknownGroupIdError) [Error](https://golang.org/src/os/user/user.go?s=2472:2515#L70)
<pre>func (e <a href="#UnknownGroupIdError">UnknownGroupIdError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnknownUserError">type</a> [UnknownUserError](https://golang.org/src/os/user/user.go?s=2236:2264#L60)
UnknownUserError is returned by Lookup when
a user cannot be found.


<pre>type UnknownUserError <a href="/pkg/builtin/#string">string</a></pre>











### <a id="UnknownUserError.Error">func</a> (UnknownUserError) [Error](https://golang.org/src/os/user/user.go?s=2266:2306#L62)
<pre>func (e <a href="#UnknownUserError">UnknownUserError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="UnknownUserIdError">type</a> [UnknownUserIdError](https://golang.org/src/os/user/user.go?s=2030:2057#L52)
UnknownUserIdError is returned by LookupId when a user cannot be found.


<pre>type UnknownUserIdError <a href="/pkg/builtin/#int">int</a></pre>











### <a id="UnknownUserIdError.Error">func</a> (UnknownUserIdError) [Error](https://golang.org/src/os/user/user.go?s=2059:2101#L54)
<pre>func (e <a href="#UnknownUserIdError">UnknownUserIdError</a>) Error() <a href="/pkg/builtin/#string">string</a></pre>



## <a id="User">type</a> [User](https://golang.org/src/os/user/user.go?s=893:1757#L19)
User represents a user account.


<pre>type User struct {
<span id="User.Uid"></span>    <span class="comment">// Uid is the user ID.</span>
    <span class="comment">// On POSIX systems, this is a decimal number representing the uid.</span>
    <span class="comment">// On Windows, this is a security identifier (SID) in a string format.</span>
    <span class="comment">// On Plan 9, this is the contents of /dev/user.</span>
    Uid <a href="/pkg/builtin/#string">string</a>
<span id="User.Gid"></span>    <span class="comment">// Gid is the primary group ID.</span>
    <span class="comment">// On POSIX systems, this is a decimal number representing the gid.</span>
    <span class="comment">// On Windows, this is a SID in a string format.</span>
    <span class="comment">// On Plan 9, this is the contents of /dev/user.</span>
    Gid <a href="/pkg/builtin/#string">string</a>
<span id="User.Username"></span>    <span class="comment">// Username is the login name.</span>
    Username <a href="/pkg/builtin/#string">string</a>
<span id="User.Name"></span>    <span class="comment">// Name is the user&#39;s real or display name.</span>
    <span class="comment">// It might be blank.</span>
    <span class="comment">// On POSIX systems, this is the first (or only) entry in the GECOS field</span>
    <span class="comment">// list.</span>
    <span class="comment">// On Windows, this is the user&#39;s display name.</span>
    <span class="comment">// On Plan 9, this is the contents of /dev/user.</span>
    Name <a href="/pkg/builtin/#string">string</a>
<span id="User.HomeDir"></span>    <span class="comment">// HomeDir is the path to the user&#39;s home directory (if they have one).</span>
    HomeDir <a href="/pkg/builtin/#string">string</a>
}
</pre>









### <a id="Current">func</a> [Current](https://golang.org/src/os/user/lookup.go?s=390:419#L4)
<pre>func Current() (*<a href="#User">User</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Current returns the current user.

The first call will cache the current user information.
Subsequent calls will return the cached value and will not reflect
changes to the current user.




### <a id="Lookup">func</a> [Lookup](https://golang.org/src/os/user/lookup.go?s=770:813#L22)
<pre>func Lookup(username <a href="/pkg/builtin/#string">string</a>) (*<a href="#User">User</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
Lookup looks up a user by username. If the user cannot be found, the
returned error is of type UnknownUserError.




### <a id="LookupId">func</a> [LookupId](https://golang.org/src/os/user/lookup.go?s=1052:1092#L31)
<pre>func LookupId(uid <a href="/pkg/builtin/#string">string</a>) (*<a href="#User">User</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
LookupId looks up a user by userid. If the user cannot be found, the
returned error is of type UnknownUserIdError.






### <a id="User.GroupIds">func</a> (\*User) [GroupIds](https://golang.org/src/os/user/lookup.go?s=1678:1721#L51)
<pre>func (u *<a href="#User">User</a>) GroupIds() ([]<a href="/pkg/builtin/#string">string</a>, <a href="/pkg/builtin/#error">error</a>)</pre>
GroupIds returns the list of group IDs that the user is a member of.








