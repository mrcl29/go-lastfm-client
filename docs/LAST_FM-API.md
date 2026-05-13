# Last.fm Music Discovery API

The Last.fm API allows anyone to build their own programs using Last.fm data. Find out more about how you can plug directly into our vast database or browse the list of methods on the left.

## Table of Contents

- [Last.fm Music Discovery API](#lastfm-music-discovery-api)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [API Guides](#api-guides)
    - [Introduction](#introduction)
      - [Encoding](#encoding)
      - [Request Styles](#request-styles)
      - [Authentication](#authentication)
      - [Scrobbling](#scrobbling)
      - [Discussion](#discussion)
      - [Terms of Service](#terms-of-service)
    - [User Authentication](#user-authentication)
      - [Authentication: Introduction](#authentication-introduction)
        - [1. Get API Key](#1-get-api-key)
        - [2. Configure Your API Account](#2-configure-your-api-account)
        - [3. Choose your authentication path](#3-choose-your-authentication-path)
        - [4. Authentication Spec](#4-authentication-spec)
      - [Authentication: Web Application How-To](#authentication-web-application-how-to)
        - [1. Get an API Key](#1-get-an-api-key)
        - [2. Request authorization from the user](#2-request-authorization-from-the-user)
          - [2.1 Custom callback url](#21-custom-callback-url)
        - [3. Create an authentication handler](#3-create-an-authentication-handler)
          - [3.1 Authentication Tokens](#31-authentication-tokens)
        - [4. Fetch a Web Service Session](#4-fetch-a-web-service-session)
          - [4.1 Session Lifetime](#41-session-lifetime)
        - [5. Make authenticated web service calls](#5-make-authenticated-web-service-calls)
        - [6. Sign your calls](#6-sign-your-calls)
      - [Authentication: Mobile Application How-To](#authentication-mobile-application-how-to)
        - [1. Get an API Key](#1-get-an-api-key-1)
        - [2. Request authorization from the user](#2-request-authorization-from-the-user-1)
          - [2.1 Session Lifetime](#21-session-lifetime)
        - [3. Make authenticated web service calls](#3-make-authenticated-web-service-calls)
        - [4. Sign your calls](#4-sign-your-calls)
      - [Authentication: Desktop Application How-To](#authentication-desktop-application-how-to)
        - [1. Get an API Key](#1-get-an-api-key-2)
        - [2. Fetch a request token](#2-fetch-a-request-token)
        - [3. Request authorization from the user](#3-request-authorization-from-the-user)
        - [4. Fetch A Web Service Session](#4-fetch-a-web-service-session-1)
          - [4.1 Session Lifetime](#41-session-lifetime-1)
        - [5. Make authenticated web service calls](#5-make-authenticated-web-service-calls-1)
        - [6. Sign your calls](#6-sign-your-calls-1)
      - [Authentication API](#authentication-api)
        - [1. Authors](#1-authors)
        - [2. Requirements](#2-requirements)
          - [2.1 Web-based Authentication](#21-web-based-authentication)
        - [3. Authentication For Web Applications](#3-authentication-for-web-applications)
          - [3.1 Request authorization from the user](#31-request-authorization-from-the-user)
          - [3.2 Create an authentication handler](#32-create-an-authentication-handler)
          - [3.3 Create a Web Service Session](#33-create-a-web-service-session)
        - [4. Authentication For Desktop Applications](#4-authentication-for-desktop-applications)
          - [4.1 Fetch a request token](#41-fetch-a-request-token)
          - [4.2 Request authorization from the user](#42-request-authorization-from-the-user)
          - [4.3 Create a Web Service Session](#43-create-a-web-service-session)
        - [5. Authentication For Mobile Applications](#5-authentication-for-mobile-applications)
        - [6. Tokens \& Sessions](#6-tokens--sessions)
          - [6.1 Authentication Tokens](#61-authentication-tokens)
          - [6.2 Session Lifetime](#62-session-lifetime)
        - [7. Making authenticated calls](#7-making-authenticated-calls)
        - [8. Signing Calls](#8-signing-calls)
    - [Scrobbling 2.0 Documentation](#scrobbling-20-documentation)
      - [Overview](#overview)
      - [Now Playing Requests](#now-playing-requests)
        - [Sending a Request](#sending-a-request)
        - [How we handle requests](#how-we-handle-requests)
        - [Error handling](#error-handling)
      - [Scrobble Requests](#scrobble-requests)
        - [When is a scrobble a scrobble?](#when-is-a-scrobble-a-scrobble)
        - [When to set the "chosenByUser" parameter](#when-to-set-the-chosenbyuser-parameter)
        - [Sending a Request](#sending-a-request-1)
        - [How we handle requests](#how-we-handle-requests-1)
        - [Error handling](#error-handling-1)
        - [Retrying cached scrobbles](#retrying-cached-scrobbles)
      - [Filtered Requests](#filtered-requests)
      - [Meta data corrections](#meta-data-corrections)
      - [Help](#help)
    - [Radio API](#radio-api)
      - [Who can I stream radio to?](#who-can-i-stream-radio-to)
      - [Tuning the Radio](#tuning-the-radio)
      - [Fetching Radio Content](#fetching-radio-content)
      - [Streamer Error Codes](#streamer-error-codes)
    - [Playlists API](#playlists-api)
      - [Fetching Last.fm Playlists](#fetching-lastfm-playlists)
      - [Playback](#playback)
    - [API Tools](#api-tools)
      - [.NET](#net)
      - [Actionscript](#actionscript)
      - [C sharp](#c-sharp)
      - [Java](#java)
      - [Javascript](#javascript)
      - [Objective C](#objective-c)
      - [Perl](#perl)
      - [PHP](#php)
      - [Python](#python)
      - [Qt C++](#qt-c)
      - [Ruby](#ruby)
      - [Contributing](#contributing)
    - [REST Requests](#rest-requests)
      - [REST Responses](#rest-responses)
      - [REST Errors](#rest-errors)
      - [JSON Responses](#json-responses)
      - [Note](#note)
      - [JSON Errors](#json-errors)
      - [Example failure response:](#example-failure-response)
    - [XML-RPC](#xml-rpc)
      - [XML-RPC Requests](#xml-rpc-requests)
      - [XML-RPC Responses](#xml-rpc-responses)
  - [API Methods](#api-methods)
    - [album](#album)
      - [album.addTags](#albumaddtags)
        - [Params](#params)
        - [Auth](#auth)
        - [Sample Response](#sample-response)
        - [Errors](#errors)
      - [album.getInfo](#albumgetinfo)
        - [Example URLs](#example-urls)
        - [Params](#params-1)
        - [Auth](#auth-1)
        - [Sample Response](#sample-response-1)
        - [Attributes](#attributes)
        - [Errors](#errors-1)
      - [album.getTags](#albumgettags)
        - [Example URLs](#example-urls-1)
        - [Params](#params-2)
        - [Auth](#auth-2)
        - [Sample Response](#sample-response-2)
        - [Errors](#errors-2)
      - [album.getTopTags](#albumgettoptags)
        - [Example URLs](#example-urls-2)
        - [Params](#params-3)
        - [Auth](#auth-3)
        - [Sample Response](#sample-response-3)
        - [Attributes](#attributes-1)
        - [Errors](#errors-3)
      - [album.removeTag](#albumremovetag)
        - [Params](#params-4)
        - [Auth](#auth-4)
        - [Sample Response](#sample-response-4)
        - [Errors](#errors-4)
      - [album.search](#albumsearch)
        - [Example URLs](#example-urls-3)
        - [Params](#params-5)
        - [Auth](#auth-5)
        - [Sample Response](#sample-response-5)
        - [Errors](#errors-5)
    - [artist](#artist)
      - [artist.addTags](#artistaddtags)
        - [Params](#params-6)
        - [Auth](#auth-6)
        - [Sample Response](#sample-response-6)
        - [Errors](#errors-6)
      - [artist.getCorrection](#artistgetcorrection)
        - [Example URLs](#example-urls-4)
        - [Params](#params-7)
        - [Auth](#auth-7)
        - [Sample Response](#sample-response-7)
        - [Errors](#errors-7)
      - [artist.getInfo](#artistgetinfo)
        - [Example URLs](#example-urls-5)
        - [Params](#params-8)
        - [Auth](#auth-8)
        - [Sample Response](#sample-response-8)
        - [Errors](#errors-8)
      - [artist.getSimilar](#artistgetsimilar)
        - [Example URLs](#example-urls-6)
        - [Params](#params-9)
        - [Auth](#auth-9)
        - [Sample Response](#sample-response-9)
        - [Attributes](#attributes-2)
        - [Errors](#errors-9)
      - [artist.getTags](#artistgettags)
        - [Example URLs](#example-urls-7)
        - [Params](#params-10)
        - [Auth](#auth-10)
        - [Sample Response](#sample-response-10)
        - [Errors](#errors-10)
      - [artist.getTopAlbums](#artistgettopalbums)
        - [Example URLs](#example-urls-8)
        - [Params](#params-11)
        - [Auth](#auth-11)
        - [Sample Response](#sample-response-11)
        - [Errors](#errors-11)
      - [artist.getTopTags](#artistgettoptags)
        - [Example URLs](#example-urls-9)
        - [Params](#params-12)
        - [Auth](#auth-12)
        - [Sample Response](#sample-response-12)
        - [Errors](#errors-12)
      - [artist.getTopTracks](#artistgettoptracks)
        - [Example URLs](#example-urls-10)
        - [Params](#params-13)
        - [Auth](#auth-13)
        - [Sample Response](#sample-response-13)
        - [Errors](#errors-13)
      - [artist.removeTag](#artistremovetag)
        - [Params](#params-14)
        - [Auth](#auth-14)
        - [Sample Response](#sample-response-14)
        - [Errors](#errors-14)
      - [artist.search](#artistsearch)
        - [Example URLs](#example-urls-11)
        - [Params](#params-15)
        - [Auth](#auth-15)
        - [Sample Response](#sample-response-15)
        - [Errors](#errors-15)
    - [auth](#auth-16)
      - [auth.getMobileSession](#authgetmobilesession)
        - [Params](#params-16)
        - [Auth](#auth-17)
        - [Sample Response](#sample-response-16)
      - [auth.getSession](#authgetsession)
        - [Params](#params-17)
        - [Auth](#auth-18)
        - [Sample Response](#sample-response-17)
      - [auth.getToken](#authgettoken)
        - [Example URLs](#example-urls-12)
        - [Params](#params-18)
        - [Auth](#auth-19)
        - [Sample Response](#sample-response-18)
    - [chart](#chart)
      - [chart.getTopArtists](#chartgettopartists)
        - [Example URLs](#example-urls-13)
        - [Params](#params-19)
        - [Auth](#auth-20)
        - [Sample Response](#sample-response-19)
        - [Errors](#errors-16)
      - [chart.getTopTags](#chartgettoptags)
        - [Example URLs](#example-urls-14)
        - [Params](#params-20)
        - [Auth](#auth-21)
        - [Sample Response](#sample-response-20)
        - [Errors](#errors-17)
      - [chart.getTopTracks](#chartgettoptracks)
        - [Example URLs](#example-urls-15)
        - [Params](#params-21)
        - [Auth](#auth-22)
        - [Sample Response](#sample-response-21)
        - [Errors](#errors-18)
    - [geo](#geo)
      - [geo.getTopArtists](#geogettopartists)
        - [Example URLs](#example-urls-16)
        - [Params](#params-22)
        - [Auth](#auth-23)
        - [Sample Response](#sample-response-22)
        - [Errors](#errors-19)
      - [geo.getTopTracks](#geogettoptracks)
        - [Example URLs](#example-urls-17)
        - [Params](#params-23)
        - [Auth](#auth-24)
        - [Sample Response](#sample-response-23)
        - [Errors](#errors-20)
    - [library](#library)
      - [library.getArtists](#librarygetartists)
        - [Example URLs](#example-urls-18)
        - [Params](#params-24)
        - [Auth](#auth-25)
        - [Sample Response](#sample-response-24)
        - [Errors](#errors-21)
    - [tag](#tag)
      - [tag.getInfo](#taggetinfo)
        - [Example URLs](#example-urls-19)
        - [Params](#params-25)
        - [Auth](#auth-26)
        - [Sample Response](#sample-response-25)
        - [Attributes](#attributes-3)
        - [Errors](#errors-22)
      - [tag.getSimilar](#taggetsimilar)
        - [Example URLs](#example-urls-20)
        - [Params](#params-26)
        - [Auth](#auth-27)
        - [Sample Response](#sample-response-26)
        - [Errors](#errors-23)
      - [tag.getTopAlbums](#taggettopalbums)
        - [Example URLs](#example-urls-21)
        - [Params](#params-27)
        - [Auth](#auth-28)
        - [Sample Response](#sample-response-27)
        - [Errors](#errors-24)
      - [tag.getTopArtists](#taggettopartists)
        - [Example URLs](#example-urls-22)
        - [Params](#params-28)
        - [Auth](#auth-29)
        - [Sample Response](#sample-response-28)
        - [Errors](#errors-25)
      - [tag.getTopTags](#taggettoptags)
        - [Example URLs](#example-urls-23)
        - [Params](#params-29)
        - [Auth](#auth-30)
        - [Sample Response](#sample-response-29)
        - [Errors](#errors-26)
      - [tag.getTopTracks](#taggettoptracks)
        - [Example URLs](#example-urls-24)
        - [Params](#params-30)
        - [Auth](#auth-31)
        - [Sample Response](#sample-response-30)
        - [Errors](#errors-27)
      - [tag.getWeeklyChartList](#taggetweeklychartlist)
        - [Example URLs](#example-urls-25)
        - [Params](#params-31)
        - [Auth](#auth-32)
        - [Sample Response](#sample-response-31)
        - [Errors](#errors-28)
    - [track](#track)
      - [track.addTags](#trackaddtags)
        - [Params](#params-32)
        - [Auth](#auth-33)
        - [Sample Response](#sample-response-32)
        - [Errors](#errors-29)
      - [track.getCorrection](#trackgetcorrection)
        - [Example URLs](#example-urls-26)
        - [Params](#params-33)
        - [Auth](#auth-34)
        - [Sample Response](#sample-response-33)
        - [Errors](#errors-30)
      - [track.getInfo](#trackgetinfo)
        - [Example URLs](#example-urls-27)
        - [Params](#params-34)
        - [Auth](#auth-35)
        - [Sample Response](#sample-response-34)
        - [Attributes](#attributes-4)
        - [Errors](#errors-31)
      - [track.getSimilar](#trackgetsimilar)
        - [Example URLs](#example-urls-28)
        - [Params](#params-35)
        - [Auth](#auth-36)
        - [Sample Response](#sample-response-35)
        - [Errors](#errors-32)
      - [track.getTags](#trackgettags)
        - [Example URLs](#example-urls-29)
        - [Params](#params-36)
        - [Auth](#auth-37)
        - [Sample Response](#sample-response-36)
        - [Errors](#errors-33)
      - [track.getTopTags](#trackgettoptags)
        - [Example URLs](#example-urls-30)
        - [Params](#params-37)
        - [Auth](#auth-38)
        - [Sample Response](#sample-response-37)
        - [Errors](#errors-34)
      - [track.love](#tracklove)
        - [Params](#params-38)
        - [Auth](#auth-39)
        - [Sample Response](#sample-response-38)
      - [track.removeTag](#trackremovetag)
        - [Params](#params-39)
        - [Auth](#auth-40)
        - [Sample Response](#sample-response-39)
        - [Errors](#errors-35)
      - [track.scrobble](#trackscrobble)
        - [Params](#params-40)
        - [Auth](#auth-41)
        - [Sample Response](#sample-response-40)
        - [Attributes](#attributes-5)
      - [track.search](#tracksearch)
        - [Example URLs](#example-urls-31)
        - [Params](#params-41)
        - [Auth](#auth-42)
        - [Sample Response](#sample-response-41)
        - [Errors](#errors-36)
      - [track.unlove](#trackunlove)
        - [Params](#params-42)
        - [Auth](#auth-43)
        - [Sample Response](#sample-response-42)
      - [track.updateNowPlaying](#trackupdatenowplaying)
        - [Params](#params-43)
        - [Auth](#auth-44)
        - [Sample Response](#sample-response-43)
        - [Attributes](#attributes-6)
    - [user](#user)
      - [user.getFriends](#usergetfriends)
        - [Example URLs](#example-urls-32)
        - [Params](#params-44)
        - [Auth](#auth-45)
        - [Sample Response](#sample-response-44)
        - [Errors](#errors-37)
      - [user.getInfo](#usergetinfo)
        - [Example URLs](#example-urls-33)
        - [Params](#params-45)
        - [Auth](#auth-46)
        - [Sample Response](#sample-response-45)
        - [Errors](#errors-38)
      - [user.getLovedTracks](#usergetlovedtracks)
        - [Example URLs](#example-urls-34)
        - [Params](#params-46)
        - [Auth](#auth-47)
        - [Sample Response](#sample-response-46)
        - [Errors](#errors-39)
      - [user.getPersonalTags](#usergetpersonaltags)
        - [Example URLs](#example-urls-35)
        - [Params](#params-47)
        - [Auth](#auth-48)
        - [Sample Response](#sample-response-47)
        - [Errors](#errors-40)
      - [user.getRecentTracks](#usergetrecenttracks)
        - [Example URLs](#example-urls-36)
        - [Params](#params-48)
        - [Auth](#auth-49)
        - [Sample Response](#sample-response-48)
        - [Errors](#errors-41)
      - [user.getTopAlbums](#usergettopalbums)
        - [Example URLs](#example-urls-37)
        - [Params](#params-49)
        - [Auth](#auth-50)
        - [Sample Response](#sample-response-49)
        - [Errors](#errors-42)
      - [user.getTopArtists](#usergettopartists)
        - [Example URLs](#example-urls-38)
        - [Params](#params-50)
        - [Auth](#auth-51)
        - [Sample Response](#sample-response-50)
        - [Errors](#errors-43)
      - [user.getTopTags](#usergettoptags)
        - [Example URLs](#example-urls-39)
        - [Params](#params-51)
        - [Auth](#auth-52)
        - [Sample Response](#sample-response-51)
        - [Errors](#errors-44)
      - [user.getTopTracks](#usergettoptracks)
        - [Example URLs](#example-urls-40)
        - [Params](#params-52)
        - [Auth](#auth-53)
        - [Sample Response](#sample-response-52)
        - [Errors](#errors-45)
      - [user.getWeeklyAlbumChart](#usergetweeklyalbumchart)
        - [Example URLs](#example-urls-41)
        - [Params](#params-53)
        - [Auth](#auth-54)
        - [Sample Response](#sample-response-53)
        - [Errors](#errors-46)
      - [user.getWeeklyArtistChart](#usergetweeklyartistchart)
        - [Example URLs](#example-urls-42)
        - [Params](#params-54)
        - [Auth](#auth-55)
        - [Sample Response](#sample-response-54)
        - [Errors](#errors-47)
      - [user.getWeeklyChartList](#usergetweeklychartlist)
        - [Example URLs](#example-urls-43)
        - [Params](#params-55)
        - [Auth](#auth-56)
        - [Sample Response](#sample-response-55)
        - [Errors](#errors-48)
      - [user.getWeeklyTrackChart](#usergetweeklytrackchart)
        - [Example URLs](#example-urls-44)
        - [Params](#params-56)
        - [Auth](#auth-57)
        - [Sample Response](#sample-response-56)
        - [Errors](#errors-49)

## Getting Started

Our API is available to anyone. Here's what you need to get going:

- Get an API account
- Read the Documentation
- Join the Support Forums (opens new window)

> Commercial or Research Usage
> If you are planning to use our API for commercial or research/academic purposes, please contact us prior to use via email at [partners@last.fm].

## API Guides

### Introduction

The Last.fm API allows you to call methods that respond in REST (opens new window) style xml. Individual methods are detailed in the menu on the left.

> API ROOT
> The API root URL is located at [http://ws.audioscrobbler.com/2.0/](http://ws.audioscrobbler.com/2.0/) (opens new window)

Generally speaking, you will send a method parameter expressed as 'package.method' along with method specific arguments to the root URL. The API supports multiple transport formats but will respond in Last.fm idiom xml by default.

Note:

- Please use an identifiable User-Agent header on all requests. This helps our logging and reduces the risk of you getting banned.
- Be reasonable in your usage of the API and ensure you don't make an excessive number of calls as that can impact the reliability of the service to you and other users. We encourage best practice implementation, for example, if you're making a web application, try not to hit the API on page load. Your account may be suspended if your application is continuously making several calls per second or if you’re making excessive calls. See our API Terms of Service for more information on limits.
- If you are planning to use our API for commercial purposes, please contact us via email at [partners@last.fm](partners@last.fm).
- We assume that you are using an RFC 3986 (opens new window)-compliant HTTP client to access the web services. In particular, pay attention to your url encoding. This will not be an issue for 99% of developers.

#### Encoding

Use UTF-8 (opens new window) encoding when sending arguments to API methods.

#### Request Styles

You can get more information on how to work with REST requests or XML-RPC requests when calling the Last.fm API.

#### Authentication

The authentication protocol allows you to perform actions on user accounts in a manner that is secure for Last.fm users. All write services require authentication.

#### Scrobbling

We encourage services that use the Last.fm API to build-in scrobbling natively into their applications (where applicable, and particularly for media players), to allow users to send listening data in to their Last.fm user profiles. This can be done through our Scrobbling API.

#### Discussion

Join the Last.fm Support Forums (opens new window) for information about new Web Services, access to beta API's, provide feedback and discuss development with other developers.

#### Terms of Service

For our API Terms of Service please see here

### User Authentication

#### Authentication: Introduction

The authentication API provides third-parties with a secure means of creating Last.fm user sessions over the Last.fm API, for deeper integration with our platform. All write services require authentication.

##### 1. Get API Key

You will need to apply for a key before authenticating with the API.

##### 2. Configure Your API Account

Head over to your api accounts page, and select the account you wish to configure. You need to supply an application name, a description and an optional logo. Each of your account pages contains an API key and secret; you will need both to use the API.

##### 3. Choose your authentication path

- If you're building a web application, see the web application how-to for more details.
- If you're building a desktop application, see the desktop application how-to for more details.
- If you're building on a standalone device such as a mobile phone, see the mobile how-to for more details.

In some cases, you may want to choose a different authentication path from the obvious (e.g. a mobile app could well use the desktop path if there's a web browser on the device). If in doubt, check them all out.

##### 4. Authentication Spec

See the full authentication API specification for an overview of the API.

#### Authentication: Web Application How-To

This authentication how-to is for web applications only. Desktop application developers should see the desktop application how-to.

##### 1. Get an API Key

If you don’t already have an API account, please apply for one. For each of your accounts you will have a shared secret which you will require in Section 6. You will also need to set up a callback url which our authentication service will redirect to in Section 4.

##### 2. Request authorization from the user

Send your user to last.fm/api/auth with your API key as a parameter. Use an HTTP GET request. Your request will look like this:

```
http://www.last.fm/api/auth/?api_key=xxx
```

If the user is not logged in to Last.fm, they will be redirected to the login page before being asked to grant your web application permission to use their account. On this page they will see the name of your application, along with the application description and logo as supplied in Section 1.

###### 2.1 Custom callback url

You can optionally specify a callback URL that is different to your API Account callback url. Include this as a query param cb

. This allows you to have users forward to a specific part of your site after the authorisation process.

```
http://www.last.fm/api/auth/?api_key=xxx&cb=http://example.com
```

##### 3. Create an authentication handler

Once the user has granted permission to use their account on the Last.fm page, Last.fm will redirect to your callback url, supplying an authentication token as a GET variable.

```
<callback_url>/?token=xxxxxxx
```

If the callback url already contains a query string then token variable will be appended, like;

```
<callback_url>&token=xxxxxxx
```

The script located at your callback url should pick up this authentication token and use it to create a Last.fm web service session as described in Section 4.

###### 3.1 Authentication Tokens

Authentication tokens are user and API account specific. They are valid for 60 minutes from the moment they are granted.

##### 4. Fetch a Web Service Session

Send your api key along with an api signature and your authentication token as arguments to the auth.getSession API method call. The parameters for this call are defined as such:

**api_key**: Your 32-character API Key.
**token**: The authentication token received at your callback url as a GET variable.
**api_sig**: Your 32-character API method signature, as explained in Section 6

Note: You can only use an authentication token once. It will be consumed when creating your web service session.

The response format of this call is shown on the auth.getSession method page.

###### 4.1 Session Lifetime

Session keys have an infinite lifetime by default. You are recommended to store the key securely. Users are able to revoke privileges for your application on their Last.fm settings screen, rendering session keys invalid.

##### 5. Make authenticated web service calls

You can now sign your web service calls with a method signature, provided along with the session key you received in Section 4 and your API key. You will need to include all three as parameters in subsequent calls in order to be able to access services that require authentication. You can visit individual method call pages to find out if they require authentication. Your three authentication parameters are defined as:

**sk** (Required) : The session key returned by auth.getSession service.
**api_key** (Required) : Your 32-character API key.
**api_sig** (Required) : Your API method signature, constructed as explained in Section 6

##### 6. Sign your calls

Construct your api method signatures by first ordering all the parameters sent in your call alphabetically by parameter name and concatenating them into one string using a `<name><value>` scheme. So for a call to auth.getSession you may have:

```
**api_key**xxxxxxxx**method**auth.getSession**token**xxxxxxx
```

Ensure your parameters are utf8 (opens new window) encoded. Now append your secret to this string. Finally, generate an md5 (opens new window) hash of the resulting string. For example, for an account with a secret equal to 'mysecret', your api signature will be:

```
api signature = md5("api_keyxxxxxxxxmethodauth.getSessiontokenxxxxxxxmysecret")
```

Where md5() is an md5 hashing operation and its argument is the string to be hashed. The hashing operation should return a 32-character hexadecimal md5 hash.

#### Authentication: Mobile Application How-To

This authentication how-to is for standalone mobile devices only.

##### 1. Get an API Key

You can apply for an API key here. When you have been granted an API Key you can configure your accounts by visiting last.fm/api/accounts . Here you will see a shared secret which will be required in Section 4.

##### 2. Request authorization from the user

Send a request to auth.getMobileSession, sending the user's credentials to the call. The parameters for this call are defined as:

username (Required) : The last.fm username.
password (Required) : A plaintext password.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See Section 4 for more information.

This webservice has to be called via POST and HTTPS. It will fail if you try to use it via GET or HTTP.

auth.getMobileSession will return a session key in response to be used on subsequent calls.

###### 2.1 Session Lifetime

Session keys have an infinite lifetime by default. You are recommended to store the key securely. Users are able to revoke privileges for your application on their Last.fm settings screen, rendering session keys invalid.

##### 3. Make authenticated web service calls

You can now sign your web service calls with a method signature, provided along with the session key you received in Section 2 and your API key. You will need to include all three as parameters in subsequent calls in order to be able to access services that require authentication. You can visit individual method call pages to find out if they require authentication. Your three authentication parameters are defined as:

sk (Required) : The session key returned by auth.getMobileSession service.
api_key (Required) : Your 32-character API key.
api_sig (Required) : Your API method signature, constructed as explained in Section 4

##### 4. Sign your calls

Construct your api method signatures by first ordering all the parameters sent in your call alphabetically by parameter name and concatenating them into one string using a <name><value> scheme. So for a call to auth.getMobileSession you may have:

```
**api_key**xxxxxxxx**method**auth.getMobileSession**password**xxxxxxx**username**xxxxxxxx
```

Ensure your parameters are utf8 (opens new window) encoded. Now append your secret to this string. Finally, generate an md5 (opens new window) hash of the resulting string. For example, for an account with a secret equal to 'mysecret', your api signature will be:

```
api signature = md5("api_keyxxxxxxxxmethodauth.getMobileSession
                         passwordxxxxxxxusernamexxxxxxxxmysecret")
```

Where md5() is an md5 hashing operation and its argument is the string to be hashed. The hashing operation should return a 32-character hexadecimal md5 hash.

#### Authentication: Desktop Application How-To

This authentication how-to is for desktop applications only. Web application developers should see the web application how-to.

##### 1. Get an API Key

If you don’t already have an API account, please apply for one. For each of your accounts you will have a shared secret which you will require in Section 6. You will also need to set up a callback url which our authentication service will redirect to in Section 4.

##### 2. Fetch a request token

Make an API method call to the auth.getToken service. You should send the following arguments to that call:

api_key: Your 32-character API Key.
api_sig: A 32-character API method signature, constructed as explained in Section 6

This will return a token. To see the response format check the method documentation page. The token is not authorized by the user at this stage.
2.1 Authentication Tokens

Authentication tokens are API account specific. They are valid for 60 minutes from the moment they are granted.

##### 3. Request authorization from the user

Your application needs to open a web browser and send the user to last.fm/api/auth with your API key and auth token as parameters. Use an HTTP GET request. Your request will look like this:

http://www.last.fm/api/auth/?api_key=xxxxxxxxxxx&token=xxxxxxxx

If the user is not logged in to Last.fm, they will be redirected to the login page before being asked to grant your application permission to use their account. On this page they will see the name of your application, along with the application description and logo as supplied in Section 1. Once the user has granted your application permission to use their account, the browser-based process is over and the user is asked to close their browser and return to your application.

##### 4. Fetch A Web Service Session

Send your api key along with an api signature and your authentication token as arguments to the auth.getSession API method call. The parameters are defined as such:

api_key: Your 32-character API Key.
token: The authentication token received from the auth.getToken method call.
api_sig: Your 32-character API method signature, as explained in Section 6

Note: You can only use an authentication token once. It will be consumed when creating your web service session.

The response format of this call is shown on the auth.getSession method page.

###### 4.1 Session Lifetime

Session keys have an infinite lifetime by default. You are recommended to store the key securely. Users are able to revoke privileges for your application on their Last.fm settings screen, rendering session keys invalid.

##### 5. Make authenticated web service calls

You can now sign your web service calls with a method signature, provided along with the session key you received in Section 4 and your API key. You will need to include all three as parameters in subsequent calls in order to be able to access services that require authentication. You can visit individual method call pages to find out if they require authentication. Your three authentication parameters are defined as:

sk (Required) : The session key returned by auth.getSession service.
api_key (Required) : Your 32-character API key.
api_sig (Required) : Your API method signature, constructed as explained in Section 6

##### 6. Sign your calls

Construct your api method signatures by first ordering all the parameters sent in your call alphabetically by parameter name and concatenating them into one string using a <name><value> scheme. So for a call to auth.getSession you may have:

**api_key**xxxxxxxx**method**auth.getSession**token**xxxxxxx

Ensure your parameters are utf8 (opens new window) encoded. Now append your secret to this string. Finally, generate an md5 (opens new window) hash of the resulting string. For example, for an account with a secret equal to 'mysecret', your api signature will be:

api signature = md5("api_keyxxxxxxxxmethodauth.getSessiontokenxxxxxxxmysecret")

Where md5() is an md5 hashing operation and its argument is the string to be hashed. The hashing operation should return a 32-character hexadecimal md5 hash.

#### Authentication API

This is Version 1.0 of the Last.fm authentication API specification.

##### 1. Authors

Anil Bawa-Cavia (opens new window)

##### 2. Requirements

You must have applied for, and received, a Last.fm API account, via the account application screen. You must configure your Last.fm API account with:

    Your application name and description.
    Your application logo.

Your account page contains your secret which must be used when making authenticated calls – see Section 8 below.

###### 2.1 Web-based Authentication

You must also configure a callback URL which will be used in Section 3.2 below.

##### 3. Authentication For Web Applications

###### 3.1 Request authorization from the user

Web applications should send a user to last.fm/api/auth, sending an API key as a parameter, in order to authenticate the user. This should be an HTTP GET request. Your request will look like this:

http://www.last.fm/api/auth/?api_key=xxxxxxxxxx

If the user is not logged in to Last.fm, they will be redirected to the login page before being asked to grant your web application permission to use their account. On this page they will see the name of your application, along with the application description and logo as supplied in Section 2.

###### 3.2 Create an authentication handler

Once the user has granted permission to use their account on the Last.fm page, Last.fm will redirect to your callback url, supplying an authentication token as a GET variable.

<callback_url>/?token=yyyyyy

The script located at your callback url should pick up this authentication token and use it to create a Last.fm web service session as described in Section 3.3.

###### 3.3 Create a Web Service Session

Send your api key along with an api signature and your authentication token as arguments to the auth.getSession API method call. The parameters for this call are defined as such:

api_key: Your 32-character API Key.
token: The authentication token received at your callback url as a GET variable.
api_sig: Your 32-character API method signature, as explained in Section 8.

The call will respond with a session key that can be used in authenticated calls.

##### 4. Authentication For Desktop Applications

###### 4.1 Fetch a request token

Make an API method call to the auth.getToken service. You should send the following arguments to that call:

api_key: Your 32-character API Key.
api_sig: A 32-character API method signature, constructed as explained in Section 8.

This will return a token. To see the response format check the method documentation page. The token is not authorized by the user at this stage and cannot be used to create a session until it has been authorized.

###### 4.2 Request authorization from the user

Your application needs to open an instance of a web browser and send the user to last.fm/api/auth with your API key and authentication token as parameters. Use an HTTP GET request. Your request will look like this:

http://www.last.fm/api/auth/?api_key=xxxxxxxxxx&token=yyyyyy

If the user is not logged in to Last.fm, they will be redirected to the login page before being asked to grant your application permission to use their account. On this page they will see the name of your application, along with the application description and logo as supplied in Section 2. Once the user has granted your application permission to use their account, the browser-based process is over and the user is asked to close their browser and return to your application.

###### 4.3 Create a Web Service Session

Send your api key along with an api signature and your authentication token as arguments to the auth.getSession API method call. The parameters for this call are defined as such:

api_key: Your 32-character API Key.
token: The authentication token received at your callback url as a GET variable.
api_sig: Your 32-character API method signature, as explained in Section 8.

The call will respond with a session key that can be used in authenticated calls.

##### 5. Authentication For Mobile Applications

Send a request to auth.getMobileSession, sending the user's credentials to the call. The parameters for this call are defined as:

password (Required) : The user's password in plaintext.
username (Required) : The user's Last.fm username.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See Section 8 for more information.

This call must be a POST made over HTTPS.

auth.getMobileSession will return a session key in response to be used in authenticated calls.

##### 6. Tokens & Sessions

###### 6.1 Authentication Tokens

Authentication tokens are API account specific. They are valid for 60 minutes from the moment they are granted and can only used once (they are consumed when a session is created).

###### 6.2 Session Lifetime

Session keys have an infinite lifetime by default. You are recommended to store the key securely. Users are able to revoke privileges for your application on their Last.fm settings screen, rendering session keys invalid.

##### 7. Making authenticated calls

You should sign authenticated web service calls with a method signature, provided along with the session key you received from auth.getSession and your API key. You will need to include all three as parameters in authenticated calls. You can visit individual method call pages to find out if they require authentication. Your three authentication parameters are defined as:

sk (Required) : The session key returned by auth.getSession service.
api_key (Required) : Your 32-character API key.
api_sig (Required) : Your API method signature, constructed as explained in Section 8.

##### 8. Signing Calls

Sign your authenticated calls by first ordering the parameters sent in your call alphabetically by parameter name and concatenating them into one string using a <name><value> scheme. You must not include the format and callback parameters. So for a call to auth.getSession you may have:

api_keyxxxxxxxxxxmethodauth.getSessiontokenyyyyyy

Ensure your parameters are utf8 (opens new window) encoded. Now append your secret to this string. Finally, generate an md5 (opens new window) hash of the resulting string. For example, for an account with a secret equal to 'ilovecher', your api signature will be:

api signature = md5("api_keyxxxxxxxxxxmethodauth.getSessiontokenyyyyyyilovecher")

Where md5() is an md5 hashing operation and its argument is the string to be hashed. The hashing operation should return a 32-character hexadecimal md5 hash.

### Scrobbling 2.0 Documentation

#### Overview

This is a guide on how to send scrobbles to Last.fm.

Scrobbling is a way to send information about the music a user is listening to. A client is anything that plays music, such as desktop music players, mobile apps, websites, etc.

For every track a user listens to the client should send a track.updateNowPlaying request and a track.scrobble request.

Scrobbling 2.0 is not backwards compatible with the old Submissions Protocol 1.2.1 (which is deprecated).

#### Now Playing Requests

The "Now Playing" service lets a client notify Last.fm that a user has started listening to a track. This does not affect a user's charts, but will feature the current track on their profile page, along with an indication of what music player they're using.

This API method call is optional for scrobbling clients, but recommended. Requests should be sent as soon as a user starts listening to a track.

##### Sending a Request

The web service method for sending Now Playing information is track.updateNowPlaying.

As with all our write web services, requests must be sent as HTTP POST requests to http://ws.audioscrobbler.com/2.0/ with form urlencoded parameters in the body of the request. The text encoding must be UTF-8.

Requests must be authenticated.

##### How we handle requests

Once a request has been received by Last.fm the following sequence of events takes place on our side:

- Check that the request passes our filters.
- Find the track in our catalogue.
- Update the user's Now Playing status with the track and client.
- Return a response to the client indicating the outcome of the request.

##### Error handling

Last.fm signals the success or failure of a request by three different means:

- The HTTP status code.
- The lfm status attribute of the lfm XML element returned in the response body. This will be either "ok" or "failed".
- The lfm error code (when the lfm status was "failed") further describes the cause of the error. This is the error element's code attribute in the XML returned in the response body.

We recommend that your client logs all failed requests and their responses (HTTP headers, and the xml body) to assist debugging.

Now Playing requests that fail should not be retried.

#### Scrobble Requests

The scrobble service lets a client add a track-play to a user's profile. This data is used to show a user's listening history and generate personalised charts and recommendations (and more).

##### When is a scrobble a scrobble?

A track should only be scrobbled when the following conditions have been met:

- The track must be longer than 30 seconds.
- And the track has been played for at least half its duration, or for 4 minutes (whichever occurs earlier.)

As soon as these conditions have been met, the scrobble request may be sent at any time. It is often most convenient to send a scrobble request when a track has finished playing.

Other considerations:

- Do not attempt to determine a track's meta data from its filename. Please only use meta data from well-structured sources such as ID3 tags.
- Do not use the corrections returned by the now playing service as input for the scrobble request, unless they have been explicitly approved by the user.

##### When to set the "chosenByUser" parameter

This parameter is used to indicate when a scrobble comes from a source that the user doesn't have "direct" control over. In most cases where a user is scrobbling their own music you can safely ignore this parameter. However, if the user is listening to music that is effectively chosen by someone other than themselves (e.g. from a Last.fm radio stream; from some other recommendation service; or radio show put together by a DJ or host) then this value should be set to "false". If there is any ambiguity or doubt then don't send this value.

##### Sending a Request

The web service method for sending scrobbles is track.scrobble.

As with all our write web services, requests must be sent as HTTP POST requests to http://ws.audioscrobbler.com/2.0/ with form urlencoded (using utf-8) parameters in the body of the request.

Requests must be authenticated.

Multiple scrobbles may be sent in a single batch request, this is recommended when there are cached scrobbles to be sent in the case of previous errors. A batch request may contain up to 50 scrobbles.

##### How we handle requests

Once a request has been received by Last.fm the following sequence of events takes place on our side:

- Check that the request passes our filters.
- Find the track in our catalogue.
- Store the scrobble in the user's profile.
- Return a response to the client indicating the outcome of the request.

##### Error handling

Last.fm signals the success or failure of a request by three different means:

- The HTTP status code.
- The lfm status attribute of the lfm XML element returned in the response body. This will be either ok or failed.
- The lfm error code (when the lfm status was failed) further describes the cause of the error. This is the error element's code attribute in the XML returned in the response body.

We recommend that your client logs all failed requests and their responses (HTTP headers, and the xml body) to assist debugging.

No matter what the HTTP status code is, ## you must inspect the content of the response

. If the HTTP status is not 200 OK it indicates there was an error (that should be logged), but it does not indicate how to handle it. Additionally a HTTP status of 200 OK does not mean the request was successful.

For example if the request was missing a required parameter you will receive an HTTP "400 Bad Request" status and retrying without modifying the request will always give the same response.

Next inspect the lfm status and lfm status code. If the lfm status is "ok" then the request succeeded.

For example if you send a request that is missing the artist parameter the response will be the following:

```xml
<?xml version="1.0" encoding="utf-8"?>
<lfm status="failed">
    <error code="6">Missing required parameter artist</error>
</lfm>
```

Lfm error codes that indicate a scrobble request should be retried are:

- 11 : Service Offline - This service is temporarily offline, try again later.
- 16 : The service is temporarily unavailable, please try again.

Additionally this lfm error code indicates that the client should reauthenticate to get a new session key before retrying the request:

- 9 : Invalid session key - Please re-authenticate

All other error codes indicate the scrobble request was incorrectly formed in some way and should not be retried.

This diagrams describes the flow for sending scrobble requests and handling the response.

https://cdn.last.fm/images/scrobbling-error-handling.png

##### Retrying cached scrobbles

Since the server connectivity may be variable (either because of network outage, or server failure), requests will occasionally fail. It is recommended that clients hold scrobbles that need be retried in a local cache. This cache should survive client restarts, allowing the user to close the client and restart later without losing unsubmitted scrobbles. Scrobbles should be sent in order, therefore cached scrobbles should be sent before new scrobbles. Scrobbles can be sent in batches of up to 50 scrobbles per request.

#### Filtered Requests

A scrobble or Now Playing request may be ignored if we detect bad meta data. This is not treated as an error condition, so if filtering takes place the response will have an "ok" status. The server will return an ignored message with an associated ignored code. This information is useful if the client wants to show information about why a track was not added to the user's profile.

In the case of batch scrobble requests, each scrobble is filtered separately. So if only one scrobble has bad meta data and is ignored other scrobbles in the request will still be accepted.

Possible ignored message codes:

- 0 : None (the request passed all filters).
- 1 : Filtered artist.
- 2 : Filtered track.
- 3 : Timestamp too far in the past.
- 4 : Timestamp too far in the future.
- 5 : Max daily scrobbles exceeded.

We may add additional ignored codes in the future.

For example if you sent a scrobble request with artist="Unknown Artist" the response will look something like this:

```xml
<?xml version='1.0' encoding='utf-8'?>
<lfm status="ok">
    <scrobbles accepted="0" ignored="1">
        <scrobble>
            <track corrected="0">Test Track</track>
            <artist corrected="0">Unknown Artist</artist>
            <album corrected="0"></album>
            <albumartist corrected="0"></albumartist>
            <timestamp>1288728940</timestamp>
            <ignoredmessage code="1">
                Artist name failed filter: Unknown Artist
            </ignoredmessage>
        </scrobble>
    </scrobbles>
</lfm>
```

#### Meta data corrections

The Last.fm catalogue contains correction information which we use to merge mispelled artists and tracks into their correct versions. If we find such a correction when resolving a track in our catalogue we will return it in the response.

This information could optionally be used by the client to suggest track meta data (ID3 tags, etc) corrections to the user. They should not be applied automatically. Most clients will simply ignore corrections.

Corrections are indicated by a corrected="1" attribute. Its value will contain the corrected version (so the value will differ from that in the request).

For example if you send a scrobble request with artist="Bjork" the response will look something like this:

```xml
<?xml version='1.0' encoding='utf-8'?>
<lfm status="ok">
    <scrobbles accepted="1" ignored="0">
        <scrobble>
            <track corrected="0">Wanderlust</track>
            <artist corrected="1">Björk</artist>
            <album corrected="0"></album>
            <albumartist corrected="0"></albumartist>
            <timestamp>1288728745</timestamp>
            <ignoredmessage code="0"></ignoredmessage>
        </scrobble>
    </scrobbles>
</lfm>
```

This response shows the artist was corrected to "Björk".

Considerations:

- Do not use the corrections returned by the now playing service as input for the scrobble request, unless they have been explicitly approved by the user.

#### Help

If you experience any problems using the scrobbling API please report them on our support forums (opens new window). Try to give as much information about the requests and responses as possible.

### Radio API

> DEPRECATED
> This functionality is no longer supported. Documentation exists for reference only.

#### Who can I stream radio to?

> Warning
> Any API account can only stream radio to Last.fm's paid subscribers Note: Due to licensing restrictions, you may not use the radio API on mobile telephones.

#### Tuning the Radio

The radio API requires authentication. See the authentication how-to.

Once authenticated, you can tune the radio using the radio.tune API method. This takes a station parameter that must correspond to a last.fm protocol station url. Here is the general scheme of these URLs:

```
lastfm://<stationtype>/<resourcename>/<station-subtype>
```

Here is a list of the station types currently publicly available, with example protocol URLs:

- Artist
  - Similar Artists Radio (e.g. lastfm://artist/cher/similarartists )
  - Top Fans Radio (e.g. lastfm://artist/cher/fans )
- User
  - Library Radio (e.g. lastfm://user/last.hq/personal )
  - Mix Radio (e.g. lastfm://user/last.hq/loved )
  - Recommendation Radio (e.g. lastfm://user/last.hq/recommended )
  - Neighbours Radio (e.g. lastfm://user/last.hq/neighbours )
- Tag
  - Global Tag Radio (e.g. lastfm://globaltags/disco )

#### Fetching Radio Content

Once the station is tuned, use the API method radio.getPlaylist to fetch content in XSPF format (see the XSPF specification (opens new window)). You will need to periodically pull on this XSPF service as it will provide content in small chunks. We recommend you prefetch a new XSPF before reaching the end of the last.

The XSPF will look like this:

```xspf
<playlist version="1" xmlns="http://xspf.org/ns/0/">
     <title>+Cher+Similar+Artists</title>
     <creator>Last.fm</creator>
     <date>2007-11-26T17:34:38</date>
     <link rel="http://www.last.fm/expiry">3600</link>
     <trackList>
      <track>
       <location>http://play.last.fm/....mp3</location>
       <title>Two People (Live)</title>
       <identifier>8212510</identifier>
       <album>Tina Live In Europe</album>
       <creator>Tina Turner</creator>
       <duration>265000</duration>
       <image>http://images.amazon.com/images/...</image>
       <extension application="http://www.last.fm/">
        <trackauth>12345</trackauth>
        <artistpage>http://www.last.fm/music/Tina+Turner</artistpage>
        <albumpage>http://www.last.fm/music/...</albumpage>
        <trackpage>http://www.last.fm/music/...</trackpage>
        <buyTrackURL>...</buyTrackURL>
        <buyAlbumURL/>
        <freeTrackURL>
        </extension>
      </track>
      <track>
       ...
```

The expiry extension is used to communicate in seconds (from the granting of the XSPF), how long any of the track URLs in your XSPF are valid for. If you request any tracks after this expiry period you will receive errors from our streaming service (see below).

Note the extension node, which holds last.fm specific information relating to the track.

The trackauth extension is now deprecated and can be ignored.

Use the location nodes to fetch individual tracks. All tracks must be requested once and only once, in the order supplied in the XSPF. Requesting the same track multiple times will result in an HTTP error returned from our streamers. Note that the track URLs provided will force an HTTP 302 redirect to the actual track location.

Please ensure your streaming library supports HTTP redirects when fetching tracks

All tracks streamed are encoded as 128kbps MP3 (opens new window) files. We recommend clients begin playback as soon as a reasonable buffer (6-8 seconds) has been fetched. Do not attempt to store the file locally other than maintaining a reasonable buffer.

Pausing playback is strictly not allowed and our streamers will return an error if you attempt to reconnect to the track more than once.

#### Streamer Error Codes

- HTTP 403 - Invalid ticket: You may be requesting tracks in the wrong order, your playlist may have expired or you have attempted to fetch the same track URL multiple times.
- HTTP 503 - Unexpected Error: Our streamers are not healthy. Try again later.

### Playlists API

> DEPRECATED
> This functionality is no longer supported. Documentation exists for reference only.

Last.fm playlists do not contain streaming content.

#### Fetching Last.fm Playlists

Use the playlist.fetch method call to fetch XSPF (opens new window) playlists. XSPF is a web standard for sharable playlists, and XSPF libraries (opens new window) are available in several languages. playlist.fetch takes a `lastfm' protocol URL as an argument; this identifies the playlist you're requesting. The following protocol url schemes are currently supported:

- Album Playlists `lastfm://playlist/album/<album_id>`
- User Playlists `lastfm://playlist/<playlist_id>`
- Tag Playlists `lastfm://playlist/tag/<tag_name>/freetracks`

#### Playback

If you'd like to add play links to your playlist you can link to the Last.fm track page (which will be present in the XSPF extension node), appending '?autostart' to the URL.

### API Tools

> **UNOFFICIAL**
> These API tools are unofficial and are not supported by Last.fm. Please contact the individual authors with questions and patches if need be.

#### .NET

- LastFmLib.Net (opens new window)
- LPFM Last.fm Scrobbler (opens new window)

#### Actionscript

- lastfm-as3-api (opens new window)

#### C sharp

- lastfm-sharp (opens new window)

#### Java

- Last.fm API Java Bindings (opens new window)

#### Javascript

- JavaScript Last.fm API - Felix Bruns (opens new window)

#### Objective C

- FMEngine (opens new window)

#### Perl

- Net::LastFM (opens new window)

#### PHP

- PHP Last.FM API - Matt Oakes (opens new window)
    PHP last.fm API - Felix Bruns (opens new window)

#### Python

- pyLast (opens new window)

#### Qt C++

- liblastfm (opens new window)

#### Ruby

- Scrobbler2 (opens new window)

#### Contributing

If you're writing any tools that use the Last.fm API and you think others might be interested (no matter what the programming language), post a link on the support forum (opens new window). If we decide to include it as a download on this page we'll even send you some Last.fm merchandise. Because we're nice like that.

### REST Requests

> API ROOT
> The API root URL is located at [http://ws.audioscrobbler.com/2.0/](http://ws.audioscrobbler.com/2.0/) (opens new window)

Generally speaking, you will send a method parameter expressed as 'package.method' along with method specific arguments to the root URL. The following parameters are required for all calls:

**api_key** : A Last.fm API Key.
**method** : An API method expressed as package.method, corresponding to a documented last.fm API method name.

For example:

```url
http://ws.audioscrobbler.com/2.0/?method=artist.getSimilar&api_key=xxx...
```

If you are accessing a write service, you will need to submit your request as an HTTP POST request. All POST requests should be made to the root url:

```url
http://ws.audioscrobbler.com/2.0/
```

With all parameters (including the 'method') sent in the POST body. In order to perform write requests you will need to authenticate a user with the API. See authentication for more.

#### REST Responses

Responses will be wrapped in an lfm status node

```
<lfm status="$status">
    ...
</lfm>
```

Where $status is either ok or failed. If the status is failed you'll get an error code and message. You can strip the status wrapper from the response by sending a raw=true argument with your method call.

#### REST Errors

See the individual method call pages for service specific error codes. Errors will communicate a code and a message in the following format:

```
<lfm status="failed">
    <error code="10">Invalid API Key</error>
</lfm>
```

#### JSON Responses

You can request API responses in JSON format with the following parameters:

format=json : A Last.fm API Key.
callback (Optional) : A callback function name which will wrap the JSON response.

#### Note

If you don't specify a callback, there's no default, and the response will be pure JSON content with a application/json MIME type. With a callback, the MIME type is text/javascript

The response is a translation of the XML response format, converted according to the following rules:

- Attributes are expressed as string member values with the attribute name as key.
- Element child nodes are expressed as object members values with the node name as key.
- Text child nodes are expressed as string values, unless the element also contains attributes, in which case the text node is expressed as a string member value with the key #text. *
- Repeated child nodes will be grouped as an array member with the shared node name as key.

This idiom is rarely used in our XML responses.
Example success response:

```json
{
"results": {
    "tagmatches": {
      "tag": \[{
        "name": "disco",
        "count": "55483",
        "url": "www.last.fm\\/tag\\/disco"
      },
      ...
      {
        "name": "disco pop",
        "count": "160",
        "url": "www.last.fm\\/tag\\/disco%20pop"
      }\]
    },
    "for": "disco"
  }
}
```

Original XML response:

```xml
<?xml version="1.0" encoding="utf-8"?>
<lfm status="ok">
 <results for="disco" xmlns:opensearch="http://a9.com/-/spec/opensearch/1.1/">
    <opensearch:Query role="request" searchTerms="disco" startPage="1" />
    <opensearch:totalResults>2641</opensearch:totalResults>
    <opensearch:startIndex>0</opensearch:startIndex>
    <opensearch:itemsPerPage>20</opensearch:itemsPerPage>
    <tagmatches>
      <tag>
        <name>disco</name>
        <count>55483</count>
        <url>www.last.fm/tag/disco</url>
      </tag>
      ...
      <tag>
        <name>disco pop</name>
        <count>160</count>
        <url>www.last.fm/tag/disco%20pop</url>
      </tag>
    </tagmatches>
  </results>
</lfm>
```

#### JSON Errors

JSON errors do not follow the same transformation rules as success errors, but use the following simplified form:

#### Example failure response:

{
    "error": 10,
    "message": "Invalid API Key"
}

### XML-RPC

#### XML-RPC Requests

Send xml-rpc requests as HTTP POST requests to http://ws.audioscrobbler.com/2.0/ (opens new window) . Send your params as named arguments using a struct in the first param node. See the example below.

```
<methodCall>
 <methodName>user.gettoptags</methodName>
 <params>
  <param>
   <value>
    <struct>
     <member>
      <name>user</name>
      <value>
       <string>joanofarctan</string>
      </value>
     </member>
     <member>
      <name>api_key</name>
      <value>
       <string>b25b959554ed76058ac220b7b2e0a026</string>
      </value>
     </member>
    </struct>
   </value>
  </param>
 </params>
</methodCall>
```

#### XML-RPC Responses

XML-RPC requests will receive responses in XML-RPC format by default. See the XML-RPC specification (opens new window) for more information.

## API Methods

### album

#### album.addTags

Tag an album using a list of user supplied tags.

##### Params

artist (Required) : The artist name
album (Required) : The album name
tags (Required) : A comma delimited list of user supplied tags to apply to this album. Accepts a maximum of 10 tags.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.

##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.

##### Sample Response

```
<lfm status="ok">
</lfm>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### album.getInfo

Get the metadata and tracklist for an album on Last.fm using the album name or a musicbrainz id.

##### Example URLs

JSON: /2.0/?method=album.getinfo&api_key=YOUR_API_KEY&artist=Cher&album=Believe&format=json (opens new window)
XML: /2.0/?method=album.getinfo&api_key=YOUR_API_KEY&artist=Cher&album=Believe (opens new window)

##### Params

artist (Required (unless mbid)] : The artist name
album (Required (unless mbid)] : The album name
mbid (Optional) : The musicbrainz id for the album
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
username (Optional) : The username for the context of the request. If supplied, the user's playcount for this album is included in the response.
lang (Optional) : The language to return the biography in, expressed as an ISO 639 alpha-2 code.
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<album>
  <name>Believe</name>
  <artist>Cher</artist>
  <id>2026126</id>
  <mbid>61bf0388-b8a9-48f4-81d1-7eb02706dfb0</mbid>
  <url>http://www.last.fm/music/Cher/Believe</url>
  <releasedate>6 Apr 1999, 00:00</releasedate>
  <image size="small">...</image>
  <image size="medium">...</image>
  <image size="large">...</image>
  <listeners>47602</listeners>
  <playcount>212991</playcount>
  <toptags>
    <tag>
      <name>pop</name>
      <url>http://www.last.fm/tag/pop</url>
    </tag>
    ...
  </toptags>
  <tracks>
    <track rank="1">
      <name>Believe</name>
      <duration>239</duration>
      <mbid/>
      <url>http://www.last.fm/music/Cher/_/Believe</url>
      <streamable fulltrack="0">1</streamable>
      <artist>
        <name>Cher</name>
        <mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid>
        <url>http://www.last.fm/music/Cher</url>
      </artist>
    </track>
    ...
  </tracks>
</album>
```

##### Attributes

- duration : In seconds

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### album.getTags

Get the tags applied by an individual user to an album on Last.fm. To retrieve the list of top tags applied to an album by all users use album.getTopTags.

##### Example URLs

JSON: /2.0/?method=album.gettags&artist=cher&album=believe&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=album.gettags&artist=cher&album=believe&api_key=YOUR_API_KEY (opens new window)

##### Params

artist (Required (unless mbid)] : The artist name
album (Required (unless mbid)] : The album name
mbid (Optional) : The musicbrainz id for the album
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
user (Optional) : If called in non-authenticated mode you must specify the user to look up
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<tags artist="Sally Shapiro" album="Disco Romance">
  <tag>
    <name>swedish</name>
    <url>http://www.last.fm/tag/swedish</url>
  </tag>
  ...
</tags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### album.getTopTags

Get the top tags for an album on Last.fm, ordered by popularity.

##### Example URLs

JSON: /2.0/?method=album.gettoptags&artist=radiohead&album=the%20bends&api_key=YOUR_API_... (opens new window)
XML: /2.0/?method=album.gettoptags&artist=radiohead&album=the%20bends&api_key=YOUR_API_KEY (opens new window)

##### Params

artist (Required (unless mbid)] : The artist name
album (Required (unless mbid)] : The album name
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
mbid (Optional) : The musicbrainz id for the album
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<?xml version="1.0" encoding="utf-8"?>
<lfm status="ok">
    <toptags artist="Radiohead" album="The Bends">
        <tag>
            <name>albums I own</name>
            <count>100</count>
            <url>http://www.last.fm/tag/albums%20i%20own</url>
        </tag>
        ...
    </toptags>
</lfm>
```

##### Attributes

- count : A weighted count of how often the tag was applied, with a maximum of 100

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### album.removeTag

Remove a user's tag from an album.

##### Params

artist (Required) : The artist name
album (Required) : The album name
tag (Required) : A single user tag to remove from this album.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.

##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.

##### Sample Response

```
<lfm status="ok">
</lfm>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### album.search

Search for an album by name. Returns album matches sorted by relevance.

##### Example URLs

JSON: /2.0/?method=album.search&album=believe&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=album.search&album=believe&api_key=YOUR_API_KEY (opens new window)

##### Params

limit (Optional) : The number of results to fetch per page. Defaults to 30.
page (Optional) : The page number to fetch. Defaults to first page.
album (Required) : The album name
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<results for="believe">
  <opensearch:Query role="request" searchTerms="believe" startPage="1"/>
  <opensearch:totalResults>734</opensearch:totalResults>
  <opensearch:startIndex>0</opensearch:startIndex>
  <opensearch:itemsPerPage>20</opensearch:itemsPerPage>
  <albummatches>
    <album>
      <name>Make Believe</name>
      <artist>Weezer</artist>
      <id>2025180</id>
      <url>http://www.last.fm/music/Weezer/Make+Believe</url>
      <image size="small">http://userserve-ak.last.fm/serve/34/8673675.jpg</image>
      <image size="medium">http://userserve-ak.last.fm/serve/64/8673675.jpg</image>
      <image size="large">http://userserve-ak.last.fm/serve/126/8673675.jpg</image>
      <streamable>0</streamable>
    </album>
    ...
  </albummatches>
</results>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### artist

#### artist.addTags

Tag an artist with one or more user supplied tags.

##### Params

artist (Required) : The artist name
tags (Required) : A comma delimited list of user supplied tags to apply to this artist. Accepts a maximum of 10 tags.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.

##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.

##### Sample Response

```
<lfm status="ok">
</lfm>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getCorrection

Use the last.fm corrections data to check whether the supplied artist has a correction to a canonical artist

##### Example URLs

JSON: /2.0/?method=artist.getcorrection&artist=Guns and Roses&api_key=YOUR_API_KEY (opens new window)
XML: /2.0/?method=artist.getcorrection&artist=Guns and Roses&api_key=YOUR_API_KEY (opens new window)

##### Params

artist (Required) : The artist name to correct.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.

##### Sample Response

```
<corrections>
  <correction index="0">
    <artist>
      <name>Guns N' Roses</name>
      <mbid>eeb1195b-f213-4ce1-b28c-8565211f8e43</mbid>
      <url>http://www.last.fm/music/Guns+N%27+Roses</url>
    </artist>
  </correction>
</corrections>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getInfo

Get the metadata for an artist. Includes biography, truncated at 300 characters.
##### Example URLs

JSON: /2.0/?method=artist.getinfo&artist=Cher&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=artist.getinfo&artist=Cher&api_key=YOUR_API_KEY (opens new window)
##### Params

artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the artist
lang (Optional) : The language to return the biography in, expressed as an ISO 639 alpha-2 code.
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
username (Optional) : The username for the context of the request. If supplied, the user's playcount for this artist is included in the response.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<artist>
  <name>Cher</name>
  <mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid>
  <url>http://www.last.fm/music/Cher</url>
  <image size="small">http://userserve-ak.last.fm/serve/50/285717.jpg</image>
  <image size="medium">http://userserve-ak.last.fm/serve/85/285717.jpg</image>
  <image size="large">http://userserve-ak.last.fm/serve/160/285717.jpg</image>
  <streamable>1</streamable>
  <stats>
    <listeners>196440</listeners>
    <plays>1599101</plays>
  </stats>
  <similar>
    <artist>
      <name>Madonna</name>
      <url>http://www.last.fm/music/Madonna</url>
      <image size="small">http://userserve-ak.last.fm/serve/50/5112299.jpg</image>
      <image size="medium">http://userserve-ak.last.fm/serve/85/5112299.jpg></image>
      <image size="large">http://userserve-ak.last.fm/serve/160/5112299.jpg</image>
    </artist>
    ...
  </similar>
  <tags>
    <tag>
      <name>pop</name>
      <url>http://www.last.fm/tag/pop</url>
    </tag>
    ...
  </tags>
  <bio>
    <published>Thu, 13 Mar 2008 03:59:18 +0000</published>
    <summary>...</summary>
    <content>...</content>
  </bio>
</artist>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getSimilar

Get all the artists similar to this artist
##### Example URLs

JSON: /2.0/?method=artist.getsimilar&artist=cher&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=artist.getsimilar&artist=cher&api_key=YOUR_API_KEY (opens new window)
##### Params

limit (Optional) : Limit the number of similar artists returned
artist (Required (unless mbid)] : The artist name
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
mbid (Optional) : The musicbrainz id for the artist
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<similarartists artist="Cher">
  <artist>
    <name>Sonny & Cher</name>
    <mbid>3d6e4b6d-2700-458c-9722-9021965a8164</mbid>
    <match>1</match>
    <url>www.last.fm/music/Sonny%2B%2526%2BCher</url>
    <image size="small">http://userserve-ak.last.fm/serve/34/71168880.png</image>
    <image size="medium">http://userserve-ak.last.fm/serve/64/71168880.png</image>
    <image size="large">http://userserve-ak.last.fm/serve/126/71168880.png</image>
    <image size="extralarge">http://userserve-ak.last.fm/serve/252/71168880.png</image>
    <image size="mega">http://userserve-ak.last.fm/serve/500/71168880/Sonny++Cher.png</image>
    <streamable>1</streamable>
  </artist>
  ...
</similarartists>
```

##### Attributes

    match : A similarity value between 0 (not similar) and 1 (very similar)

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getTags

Get the tags applied by an individual user to an artist on Last.fm. If accessed as an authenticated service /and/ you don't supply a user parameter then this service will return tags for the authenticated user. To retrieve the list of top tags applied to an artist by all users use artist.getTopTags.

##### Example URLs

JSON: /2.0/?method=artist.getTags&artist=Red%20Hot%20Chili%20Peppers&user=RJ&api_key=YOU... (opens new window)
XML: /2.0/?method=artist.getTags&artist=Red%20Hot%20Chili%20Peppers&user=RJ&api_key=YOU... (opens new window)

##### Params

artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the artist
user (Optional) : If called in non-authenticated mode you must specify the user to look up
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<tags artist="Sally Shapiro">
  <tag>
    <name>italo</name>
    <url>http://www.last.fm/tag/italo</url>
  </tag>
  ...
</tags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getTopAlbums

Get the top albums for an artist on Last.fm, ordered by popularity.
##### Example URLs

JSON: /2.0/?method=artist.gettopalbums&artist=cher&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=artist.gettopalbums&artist=cher&api_key=YOUR_API_KEY (opens new window)
##### Params

artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the artist
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
page (Optional) : The page number to fetch. Defaults to first page.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<topalbums artist="Cher">
  <album rank="1">
    <name>Believe</name>
    <mbid>61bf0388-b8a9-48f4-81d1-7eb02706dfb0</mbid>
    <listeners>24486</listeners>
    <url>http://www.last.fm/music/Cher/Believe</url>
    <image size="small">...</image>
    <image size=" medium">...</image>
    <image size="large">...</image>
  </album>
  ...
</topalbums>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getTopTags

Get the top tags for an artist on Last.fm, ordered by popularity.
##### Example URLs

JSON: /2.0/?method=artist.gettoptags&artist=cher&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=artist.gettoptags&artist=cher&api_key=YOUR_API_KEY (opens new window)
##### Params

artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the artist
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptags artist="Cher">
  <tag>
    <name>pop</name>
    <url>http://www.last.fm/tag/pop</url>
  </tag>
  ...
</toptags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.getTopTracks

Get the top tracks by an artist on Last.fm, ordered by popularity
##### Example URLs

JSON: /2.0/?method=artist.gettoptracks&artist=cher&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=artist.gettoptracks&artist=cher&api_key=YOUR_API_KEY (opens new window)
##### Params

artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the artist
autocorrect[0|1] (Optional) : Transform misspelled artist names into correct artist names, returning the correct version instead. The corrected artist name will be returned in the response.
page (Optional) : The page number to fetch. Defaults to first page.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptracks artist="Cher">
  <track rank="1">
    <name>Believe</name>
    <mbid/>
    <playcount>56325</playcount>
    <listeners>23217</listeners>
    <url>http://www.last.fm/music/Cher/_/Believe</url>
    <image size="small">...</image>
    <image size=" medium">...</image>
    <image size="large">...</image>
  </track>
  ...
</toptracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.removeTag

Remove a user's tag from an artist.
##### Params

artist (Required) : The artist name
tag (Required) : A single user tag to remove from this artist.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response

```
<lfm status="ok">
</lfm>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### artist.search

Search for an artist by name. Returns artist matches sorted by relevance.
##### Example URLs

JSON: /2.0/?method=artist.search&artist=cher&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=artist.search&artist=cher&api_key=YOUR_API_KEY (opens new window)
##### Params

limit (Optional) : The number of results to fetch per page. Defaults to 30.
page (Optional) : The page number to fetch. Defaults to first page.
artist (Required) : The artist name
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<results for="cher" xmlns:opensearch="http://a9.com/-/spec/opensearch/1.1/">
  <opensearch:Query role="request" searchTerms="cher" startPage="1"/>
  <opensearch:totalResults>386</opensearch:totalResults>
  <opensearch:startIndex>0</opensearch:startIndex>
  <opensearch:itemsPerPage>20</opensearch:itemsPerPage>
  <artistmatches>
    <artist>
      <name>Cher</name>
      <mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid>
      <url>www.last.fm/music/Cher</url>
      <image_small>http://userserve-ak.last.fm/serve/50/342437.jpg</image_small>
      <image>http://userserve-ak.last.fm/serve/160/342437.jpg</image>
      <streamable>1</streamable>
    </artist>
	...
  </artistmatches>
</results>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### auth

#### auth.getMobileSession

Create a web service session for a user. Used for authenticating a user when the password can be inputted by the user. Accepts email address as well, so please use the username supplied in the output. Only suitable for standalone mobile devices. See the authentication how-to for more. You must use HTTPS and POST in order to use this method.
##### Params

password (Required) : The password in plain text.
username (Required) : The last.fm username or email address.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
##### Auth

This service does not require authentication.
##### Sample Response

```
<lfm status="ok">
  <session>
    <name>MyLastFMUsername</name>
    <key>d580d57f32848f5dcf574d1ce18d78b2</key>
    <subscriber>0</subscriber>
  </session>
</lfm>
```

DEPRECATED BEHAVIOUR

This method has other parameters which are now deprecated and should not be used.

A previous version of auth.getMobileSession accepted an authToken parameter in place of password, where authToken was defined as:

authToken = md5(username + md5(password))

We recommend all clients update to the latest version as support for authToken will be removed in the future.
Errors

    4 : You must use POST in order to use this method.
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### auth.getSession

Fetch a session key for a user. The third step in the authentication process. See the authentication how-to for more information.
##### Params

token (Required) : A 32-character ASCII hexadecimal MD5 hash returned by step 1 of the authentication process (following the granting of permissions to the application by the user)
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
##### Auth

This service does not require authentication.
##### Sample Response

```
<lfm status="ok">
  <session>
    <name>MyLastFMUsername</name>
    <key>d580d57f32848f5dcf574d1ce18d78b2</key>
     <subscriber>0</subscriber>
  </session>
</lfm>
```

Errors

- 4 : Invalid authentication token supplied
- 14 : This token has not been authorized
- 15 : This token has expired
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### auth.getToken

Fetch an unathorized request token for an API account. This is step 2 of the authentication process for desktop applications. Web applications do not need to use this service.
##### Example URLs

JSON: /2.0/?method=auth.gettoken&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=auth.gettoken&api_key=YOUR_API_KEY (opens new window)
##### Params

api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
##### Auth

This service does not require authentication.
##### Sample Response

```
<lfm status="ok">
  <token>cf45fe5a3e3cebe168480a086d7fe481</token>
</lfm>
```

Errors

- 8 : There was an error granting the request token. Please try again later
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### chart

#### chart.getTopArtists

Get the top artists chart
##### Example URLs

JSON: /2.0/?method=chart.gettopartists&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=chart.gettopartists&api_key=YOUR_API_KEY (opens new window)
##### Params

page (Optional) : The page number to fetch. Defaults to first page.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<artists page="1" perPage="50" totalPages="20" total="1000">
  <artist>
    <name>The Beatles</name>
    <playcount>1550293</playcount>
    <listeners>114106</listeners>
    <mbid>b10bbbfc-cf9e-42e0-be17-e2c3e1d2600d</mbid>
    <url>http://www.last.fm/music/The+Beatles</url>
    <streamable>1</streamable>
    <image size="small">http://userserve-ak.last.fm/serve/34/880929.jpg</image>
    <image size="medium">http://userserve-ak.last.fm/serve/64/880929.jpg</image>
    <image size="large">http://userserve-ak.last.fm/serve/126/880929.jpg</image>
    <image size="extralarge">http://userserve-ak.last.fm/serve/252/880929.jpg</image>
    <image size="mega">http://userserve-ak.last.fm/serve/500/880929/The+Beatles.jpg</image>
  </artist>
  ...
</artists>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### chart.getTopTags

Get the top artists chart
##### Example URLs

JSON: /2.0/?method=chart.gettoptags&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=chart.gettoptags&api_key=YOUR_API_KEY (opens new window)
##### Params

page (Optional) : The page number to fetch. Defaults to first page.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<tags page="1" perPage="50" totalPages="5" total="250">
  <tag>
    <name>rock</name>
    <url>http://www.last.fm/tag/rock</url>
    <reach>309437</reach>
    <taggings>3064604</taggings>
    <streamable>1</streamable>
    <wiki>
      <published>Sun, 24 Oct 2010 17:40:33 +0000</published>
      <summary>
Rock music is a genre of music started in America. It h...
      </summary>
      <content>
Rock music is a genre of music started in America. It has its roots in 1940s and 1950s rock and roll and rockabilly, which evolved from blues, country music and other influences. According to the All Music Guide, “In its pu...
      </content>
    </wiki>
</tag>
...
</tags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### chart.getTopTracks

Get the top tracks chart
##### Example URLs

JSON: /2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=chart.gettoptracks&api_key=YOUR_API_KEY (opens new window)
##### Params

page (Optional) : The page number to fetch. Defaults to first page.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<tracks page="1" perPage="50" totalPages="20" total="1000">
  <track>
    <name>Dark Fantasy</name>
    <playcount>124394</playcount>
    <listeners>42141</listeners>
    <mbid/>
    <url>http://www.last.fm/music/Kanye+West/_/Dark+Fantasy</url>
    <streamable fulltrack="0">0</streamable>
    <artist>
      <name>Kanye West</name>
      <mbid>164f0d73-1234-4e2c-8743-d77bf2191051</mbid>
      <url>http://www.last.fm/music/Kanye+West</url>
    </artist>
  </track>
  ...
</tracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### geo

#### geo.getTopArtists

Get the most popular artists on Last.fm by country

##### Example URLs

JSON: /2.0/?method=geo.gettopartists&country=spain&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=geo.gettopartists&country=spain&api_key=YOUR_API_KEY (opens new window)
##### Params

country (Required) : A country name, as defined by the ISO 3166-1 country names standard
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<topartists country="Spain">
  <artist rank="1">
    <name>Coldplay</name>
    <playcount>3199</playcount>
    <mbid>cc197bad-dc9c-440d-a5b5-d52ba2e14234</mbid>
    <url>http://www.last.fm/music/Coldplay</url>
    <streamable>1</streamable>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </artist>
  ...
</topartists>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### geo.getTopTracks

Get the most popular tracks on Last.fm last week by country
##### Example URLs

JSON: /2.0/?method=geo.gettoptracks&country=spain&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=geo.gettoptracks&country=spain&api_key=YOUR_API_KEY (opens new window)
##### Params

country (Required) : A country name, as defined by the ISO 3166-1 country names standard
location (Optional) : A metro name, to fetch the charts for (must be within the country specified)
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptracks country="Spain">  
  <track rank="1">
    <name>Violet Hill</name>
    <playcount>1055</playcount>
    <mbid/>
    <url>http://www.last.fm/music/Coldplay/_/Violet+Hill</url>
    <streamable fulltrack="0">1</streamable>
    <artist>
      <name>Coldplay</name>
      <mbid>cc197bad-dc9c-440d-a5b5-d52ba2e14234</mbid>
      <url>http://www.last.fm/music/Coldplay</url>
    </artist>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </track>
  ...
</toptracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### library

#### library.getArtists

A paginated list of all the artists in a user's library, with play counts and tag counts.
##### Example URLs

JSON: /2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan&format=json (opens new window)
XML: /2.0/?method=library.getartists&api_key=YOUR_API_KEY&user=joanofarctan (opens new window)
##### Params

user (Required) : The user whose library you want to fetch.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number you wish to scan to.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<artists user="RJ" page="1" perPage="50" totalPages="20">
  <artist rank="1">
    <name>Dream Theater</name>
    <playcount>1346</playcount>
    <tagcount>0</tagcount>
    <mbid>28503ab7-8bf2-4666-a7bd-2644bfc7cb1d</mbid>
    <url>http://www.last.fm/music/Dream+Theater</url>
    <streamable>1</streamable>
    <image size="small">http://userserve-ak.last.fm/serve/50/95853.jpg</image>
    <image size="medium">http://userserve-ak.last.fm/serve/85/95853.jpg</image>
    <image size="large">http://userserve-ak.last.fm/serve/160/95853.jpg</image>
  </artist>
  ...
</artists>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### tag

#### tag.getInfo

Get the metadata for a tag
##### Example URLs

JSON: /2.0/?method=tag.getinfo&tag=disco&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.getinfo&tag=disco&api_key=YOUR_API_KEY (opens new window)
##### Params

lang (Optional) : The language to return the wiki in, expressed as an ISO 639 alpha-2 code.
tag (Required) : The tag name
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<tag>
    <name>disco</name>
    <url>http://www.last.fm/tag/disco</url>
    <reach>27199</reach>
    <taggings>114210</taggings>
    <streamable>1</streamable>
    <wiki>
        <published>Thu, 19 Aug 2010 03:22:16 +0000</published>
        <summary><!\[CDATA\[Disco is a genre of dance-oriented music\]\]></summary>
        <content><!\[CDATA\[Disco is a genre of dance-oriented music that originated in African American, gay and Hispanic American communities in 1970s. (truncated for readability in sample)\]\]></content>
    </wiki>
</tag>
```

##### Attributes

    reach : The number of users that have used this tag
    streamable : A value of '1' indicates this tag can be used as a radio station
    taggings : The total number of times this tag has been used

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### tag.getSimilar

Search for tags similar to this one. Returns tags ranked by similarity, based on listening data.
##### Example URLs

JSON: /2.0/?method=tag.getsimilar&tag=disco&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.getsimilar&tag=disco&api_key=YOUR_API_KEY (opens new window)
##### Params

tag (Required) : The tag name
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<similartags tag="Disco">
  <tag>
    <name>high energy</name>
    <url>http://www.last.fm/tag/high energy</url>
    <streamable>1</streamable>
  </tag>
  ...
</similartags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### tag.getTopAlbums

Get the top albums tagged by this tag, ordered by tag count.
##### Example URLs

JSON: /2.0/?method=tag.gettopalbums&tag=disco&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.gettopalbums&tag=disco&api_key=YOUR_API_KEY (opens new window)
##### Params

tag (Required) : The tag name
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<topalbums tag="Disco">
  <album rank="">
    <name>Overpowered</name>
    <mbid/>
    <url>
      http://www.last.fm/music/Róisín+Murphy/Overpowered
    </url>
    <artist>
      <name>Róisín Murphy</name>
      <mbid>4c56405d-ba8e-4283-99c3-1dc95bdd50e7</mbid>
      <url>http://www.last.fm/music/Róisín+Murphy</url>
    </artist>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </album>
  ...
</topalbums>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### tag.getTopArtists

Get the top artists tagged by this tag, ordered by tag count.
##### Example URLs

JSON: /2.0/?method=tag.gettopartists&tag=disco&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.gettopartists&tag=disco&api_key=YOUR_API_KEY (opens new window)
##### Params

tag (Required) : The tag name
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<topartists tag="Disco">
  <artist rank="">
    <name>ABBA</name>
    <mbid>d87e52c5-bb8d-4da8-b941-9f4928627dc8</mbid>
    <url>http://www.last.fm/music/ABBA</url>
    <streamable>1</streamable>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </artist>
  ...
</topartists>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### tag.getTopTags

Fetches the top global tags on Last.fm, sorted by popularity (number of times used)
##### Example URLs

JSON: /2.0/?method=tag.getTopTags&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.getTopTags&api_key=YOUR_API_KEY (opens new window)
##### Params

api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptags>
  <tag>
    <name>rock</name>
    <count>1994155</count>
    <url>www.last.fm/tag/rock</url>
  </tag>
  ...
</toptags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### tag.getTopTracks

Get the top tracks tagged by this tag, ordered by tag count.
##### Example URLs

JSON: /2.0/?method=tag.gettoptracks&tag=disco&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.gettoptracks&tag=disco&api_key=YOUR_API_KEY (opens new window)
##### Params

tag (Required) : The tag name
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptracks tag="Disco">
  <track rank="">
    <name>Stayin' Alive</name>
    <mbid/>
    <url>
      http://www.last.fm/music/Bee+Gees/_/Stayin'+Alive
    </url>
    <streamable fulltrack="0">1</streamable>
    <artist>
      <name>Bee Gees</name>
      <mbid>bf0f7e29-dfe1-416c-b5c6-f9ebc19ea810</mbid>
      <url>http://www.last.fm/music/Bee+Gees</url>
    </artist>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </track>
  ...
</toptracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### tag.getWeeklyChartList

Get a list of available charts for this tag, expressed as date ranges which can be sent to the chart services.
##### Example URLs

JSON: /2.0/?method=tag.getweeklychartlist&tag=disco&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=tag.getweeklychartlist&tag=disco&api_key=YOUR_API_KEY (opens new window)
##### Params

tag (Required) : The tag name
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<weeklychartlist tag="rock">
  <chart from="1108296002" to="1108900802"/>
  <chart from="1108900801" to="1109505601"/>
  ...
</weeklychartlist>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### track

#### track.addTags

Tag an album using a list of user supplied tags.
##### Params

artist (Required) : The artist name
track (Required) : The track name
tags (Required) : A comma delimited list of user supplied tags to apply to this track. Accepts a maximum of 10 tags.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response

```
<lfm status="ok">
</lfm>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.getCorrection

Use the last.fm corrections data to check whether the supplied track has a correction to a canonical track
##### Example URLs

JSON: [/2.0/?method=track.getcorrection&artist=guns and...](http://ws.audioscrobbler.com/2.0/?method=track.getcorrection&artist=guns and roses&track=Mrbrownstone&api_key=YOUR_API_KEY&format=json)
XML: [/2.0/?method=track.getcorrection&artist=guns and...](http://ws.audioscrobbler.com/2.0/?method=track.getcorrection&artist=guns and roses&track=Mrbrownstone&api_key=YOUR_API_KEY)
##### Params

artist (Required) : The artist name to correct.
track (Required) : The track name to correct.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<corrections>
  <correction index="0" artistcorrected="1" trackcorrected="1">
    <track>
      <name>Mr. Brownstone</name>
      <mbid/>
      <url>www.last.fm/music/Guns+N%27+Roses/_/Mr.+Brownstone</url>
      <artist>
        <name>Guns N' Roses</name>
        <mbid>eeb1195b-f213-4ce1-b28c-8565211f8e43</mbid>
        <url>http://www.last.fm/music/Guns+N%27+Roses</url>
      </artist>
    </track>
  </correction>
</corrections>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.getInfo

Get the metadata for a track on Last.fm using the artist/track name or a musicbrainz id.
##### Example URLs

JSON: /2.0/?method=track.getInfo&api_key=YOUR_API_KEY&artist=cher&track=believe&format=json (opens new window)
XML: /2.0/?method=track.getInfo&api_key=YOUR_API_KEY&artist=cher&track=believe (opens new window)
##### Params

mbid (Optional) : The musicbrainz id for the track
track (Required (unless mbid)] : The track name
artist (Required (unless mbid)] : The artist name
username (Optional) : The username for the context of the request. If supplied, the user's playcount for this track and whether they have loved the track is included in the response.
autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<track>
  <id>1019817</id>
  <name>Believe</name>
  <mbid/>
  <url>http://www.last.fm/music/Cher/_/Believe</url>
  <duration>240000</duration>
  <streamable fulltrack="1">1</streamable>
  <listeners>69572</listeners>
  <playcount>281445</playcount>
  <artist>
    <name>Cher</name>
    <mbid>bfcc6d75-a6a5-4bc6-8282-47aec8531818</mbid>
    <url>http://www.last.fm/music/Cher</url>
  </artist>
  <album position="1">
    <artist>Cher</artist>
    <title>Believe</title>
    <mbid>61bf0388-b8a9-48f4-81d1-7eb02706dfb0</mbid>
    <url>http://www.last.fm/music/Cher/Believe</url>
    <image size="small">http://userserve-ak.last.fm/serve/34/8674593.jpg</image>
    <image size="medium">http://userserve-ak.last.fm/serve/64/8674593.jpg</image>
    <image size="large">http://userserve-ak.last.fm/serve/126/8674593.jpg</image>
  </album>
  <toptags>
    <tag>
      <name>pop</name>
      <url>http://www.last.fm/tag/pop</url>
    </tag>
    ...
  </toptags>
  <wiki>
    <published>Sun, 27 Jul 2008 15:44:58 +0000</published>
    <summary>...</summary>
    <content>...</content>
  </wiki>
</track>
```

##### Attributes

    duration : In milliseconds
    fulltrack : An attribute value of 1 indicates a full length preview is available for streaming
    streamable : A tag value of 1 indicates a 30 second preview of this song is available for streaming

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.getSimilar

Get the similar tracks for this track on Last.fm, based on listening data.
##### Example URLs

JSON: /2.0/?method=track.getsimilar&artist=cher&track=believe&api_key=YOUR_API_KEY&forma... (opens new window)
XML: /2.0/?method=track.getsimilar&artist=cher&track=believe&api_key=YOUR_API_KEY (opens new window)
##### Params

track (Required (unless mbid)] : The track name
artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the track
autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.
limit (Optional) : Maximum number of similar tracks to return
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<similartracks track="Believe" artist="Cher">
  <track>
    <name>Ray of Light</name>
    <mbid/>
    <match>10.95</match>
    <url>http://www.last.fm/music/Madonna/_/Ray+of+Light</url>
    <streamable fulltrack="0">1</streamable>
    <artist>
      <name>Madonna</name>
      <mbid>79239441-bfd5-4981-a70c-55c3f15c1287</mbid>
      <url>http://www.last.fm/music/Madonna</url>
    </artist>
    <image size="small">http://cdn.last.fm/coverart/50x50/1934.jpg</image>
    <image size="medium">http://cdn.last.fm/coverart/130x130/1934.jpg</image>
    <image size="large">http://cdn.last.fm/coverart/130x130/1934.jpg</image>
  </track>
  ...
</similartracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.getTags

track.getTags

Get the tags applied by an individual user to a track on Last.fm. To retrieve the list of top tags applied to a track by all users use track.getTopTags.
##### Example URLs

JSON: /2.0/?method=track.getTags&api_key=YOUR_API_KEY&artist=AC/DC&track=Hells+Bells&use... (opens new window)
XML: /2.0/?method=track.getTags&api_key=YOUR_API_KEY&artist=AC/DC&track=Hells+Bells&use... (opens new window)
##### Params

artist (Required (unless mbid)] : The artist name
track (Required (unless mbid)] : The track name
mbid (Optional) : The musicbrainz id for the track
autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.
user (Optional) : If called in non-authenticated mode you must specify the user to look up
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<tags artist="Sally Shapiro" track="I'll be by your side">
  <tag>
    <name>swedish</name>
    <url>http://www.last.fm/tag/swedish</url>
  </tag>
  ...
</tags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.getTopTags

Get the top tags for this track on Last.fm, ordered by tag count. Supply either track & artist name or mbid.
##### Example URLs

JSON: /2.0/?method=track.gettoptags&artist=radiohead&track=paranoid+android&api_key=YOUR... (opens new window)
XML: /2.0/?method=track.gettoptags&artist=radiohead&track=paranoid+android&api_key=YOUR... (opens new window)
##### Params

track (Required (unless mbid)] : The track name
artist (Required (unless mbid)] : The artist name
mbid (Optional) : The musicbrainz id for the track
autocorrect[0|1] (Optional) : Transform misspelled artist and track names into correct artist and track names, returning the correct version instead. The corrected artist and track name will be returned in the response.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptags artist="Cher" track="Believe">
  <tag>
    <name>pop</name>
    <count>97</count>
    <url>www.last.fm/tag/pop</url>
  </tag>
  <tag>
    <name>dance</name>
    <count>88</count>
    <url>www.last.fm/tag/dance</url>
  </tag>
  ...
</toptags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.love

Love a track for a user profile.
##### Params

track (Required) : A track name (utf8 encoded)
artist (Required) : An artist name (utf8 encoded)
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response

```
<lfm status="ok">
</lfm>
```

Errors

- 8 : Operation Failed - Server Error. Please try again later.
- 6 : Invalid track name supplied
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.removeTag

Remove a user's tag from a track.
##### Params

artist (Required) : The artist name
track (Required) : The track name
tag (Required) : A single user tag to remove from this track.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response

```
<lfm status="ok">
</lfm>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.scrobble

Used to add a track-play to a user's profile. Scrobble a track, or a batch of tracks. Tracks are passed to the service using array notation for each of the below params, up to a maximum of 50 scrobbles per batch [0<=i<=49]. If you are only sending a single scrobble the array notation may be ommited. Note: Extra care should be taken while calculating the signature when using array notation as the parameter names MUST be sorted according to the ASCII table (i.e., artist[10] comes before artist[1]). It is important to not use the corrections returned by the now playing service as input for the scrobble request, unless they have been explicitly approved by the user. Parameter names are case sensitive.
##### Params

artist[i] (Required) : The artist name.
track[i] (Required) : The track name.
timestamp[i] (Required) : The time the track started playing, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
album[i] (Optional) : The album name.
context[i] (Optional) : Sub-client version (not public, only enabled for certain API keys)
streamId[i] (Optional) : The stream id for this track received from the radio.getPlaylist service, if scrobbling Last.fm radio
chosenByUser[i] (Optional) : Set to 1 if the user chose this song, or 0 if the song was chosen by someone else (such as a radio station or recommendation service). Assumes 1 if not specified
trackNumber[i] (Optional) : The track number of the track on the album.
mbid[i] (Optional) : The MusicBrainz Track ID.
albumArtist[i] (Optional) : The album artist - if this differs from the track artist.
duration[i] (Optional) : The length of the track in seconds.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response

Example for a single scrobble sent:

```
<?xml version='1.0' encoding='utf-8'?>
<lfm status="ok">
  <scrobbles accepted="1" ignored="0">
    <scrobble>
      <track corrected="0">Test Track</track>
      <artist corrected="0">Test Artist</artist>
      <album corrected="0"></album>
      <albumArtist corrected="0"></albumArtist>
      <timestamp>1287140447</timestamp>
      <ignoredMessage code="0"></ignoredMessage>
    </scrobble>
  </scrobbles>
</lfm>
```

Example for 2 scrobbles sent:

```
<?xml version='1.0' encoding='utf-8'?>
<lfm status="ok">
  <scrobbles accepted="2" ignored="0">
    <scrobble>
      <track corrected="0">Test Track 0</track>
      <artist corrected="0">Test Artist 0</artist>
      <album corrected="0"></album>
      <albumArtist corrected="0"></albumArtist>
      <timestamp>1287141093</timestamp>
      <ignoredMessage code="0"></ignoredMessage>
    </scrobble>
    <scrobble>
      <track corrected="0">Test Track 1</track>
      <artist corrected="0">Test Artist 1</artist>
      <album corrected="0"></album>
      <albumArtist corrected="0"></albumArtist>
      <timestamp>1287141093</timestamp>
      <ignoredMessage code="0"></ignoredMessage>
    </scrobble>
  </scrobbles>
</lfm>
```

##### Attributes

    accepted : Number of accepted scrobbles
    code : Ignored message codes:
        1: Artist was ignored
        2: Track was ignored
        3: Timestamp was too old
        4: Timestamp was too new
        5: Daily scrobble limit exceeded
    corrected : '1', if this track, artist or album name was automatically corrected, '0' otherwise
    ignored : Number of ignored scrobbles (see ignoredMessage for details)

Errors

- 16 : The service is temporarily unavailable, please try again.
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.search

Search for a track by track name. Returns track matches sorted by relevance.
##### Example URLs

JSON: /2.0/?method=track.search&track=Believe&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=track.search&track=Believe&api_key=YOUR_API_KEY (opens new window)
##### Params

limit (Optional) : The number of results to fetch per page. Defaults to 30.
page (Optional) : The page number to fetch. Defaults to first page.
track (Required) : The track name
artist (Optional) : Narrow your search by specifying an artist.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<results for="Believe" xmlns:opensearch="http://a9.com/-/spec/opensearch/1.1/">
  <opensearch:Query role="request" searchTerms="Believe" startPage="1"/>
  <opensearch:totalResults>25329</opensearch:totalResults>
  <opensearch:startIndex>0</opensearch:startIndex>
  <opensearch:itemsPerPage>20</opensearch:itemsPerPage>
  <trackmatches>
    <track>
      <name>Believe</name>
      <artist>Disturbed</artist>
      <url>http://www.last.fm/music/Disturbed/_/Believe</url>
      <streamable fulltrack="0">1</streamable>
      <listeners>66068</listeners>
      <image size="small">...</image>
    </track>
    ...
  </trackmatches>
</results>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.unlove

UnLove a track for a user profile.
##### Params

track (Required) : A track name (utf8 encoded)
artist (Required) : An artist name (utf8 encoded)
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response
Errors

- 8 : Operation Failed - Server Error. Please try again later.
- 6 : Invalid track name supplied
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### track.updateNowPlaying

Used to notify Last.fm that a user has started listening to a track. Parameter names are case sensitive.
##### Params

artist (Required) : The artist name.
track (Required) : The track name.
album (Optional) : The album name.
trackNumber (Optional) : The track number of the track on the album.
context (Optional) : Sub-client version (not public, only enabled for certain API keys)
mbid (Optional) : The MusicBrainz Track ID.
duration (Optional) : The length of the track in seconds.
albumArtist (Optional) : The album artist - if this differs from the track artist.
api_key (Required) : A Last.fm API key.
api_sig (Required) : A Last.fm method signature. See authentication for more information.
sk (Required) : A session key generated by authenticating a user via the authentication protocol.
##### Auth

This service requires authentication. Please see our authentication how-to.
This is a write service and must be accessed with an HTTP POST request. All parameters should be sent in the POST body, including the 'method' parameter. See rest requests for more information.
##### Sample Response

```
<?xml version='1.0' encoding='utf-8'?>
<lfm status="ok">
  <nowplaying>
    <track corrected="0">Test Track</track>
    <artist corrected="0">Test Artist</artist>
    <album corrected="0"></album>
    <albumArtist corrected="0"></albumArtist>
    <ignoredMessage code="0"></ignoredMessage>
  </nowplaying>
</lfm>
```

##### Attributes

- code : Ignored message codes:
    - 1: Artist was ignored
    - 2: Track was ignored
    - 3: Timestamp was too old
    - 4: Timestamp was too new
    - 5: Daily scrobble limit exceeded
- corrected : '1', if this track, artist or album name was automatically corrected, '0' otherwise

Errors

- 16 : The service is temporarily unavailable, please try again.
- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

### user

#### user.getFriends

Get a list of the user's friends on Last.fm.
##### Example URLs

JSON: /2.0/?method=user.getfriends&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getfriends&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The last.fm username to fetch the friends of.
recenttracks (Optional) : Whether or not to include information about friends' recent listening in the response.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

<lfm status="ok" total="109" page="1" perPage="50" totalPages="3">
    <friends user="joanofarctan">
        <user>
            <name>eartle</name>
            <realname>Michael Coffey</realname>
            <image size="small">http://userserve-ak.last.fm/serve/34/45718509.jpg</image>
            <image size="medium">http://userserve-ak.last.fm/serve/64/45718509.jpg</image>
            <image size="large">http://userserve-ak.last.fm/serve/126/45718509.jpg</image>
            <image size="extralarge">http://userserve-ak.last.fm/serve/252/45718509.jpg</image>
            <url>http://www.last.fm/user/eartle</url>
            <id>7737850</id>
            <country>UK</country>
            <age>29</age>
            <gender>m</gender>
            <subscriber>1</subscriber>
            <playcount>45366</playcount>
            <playlists>4</playlists>
            <bootstrap>0</bootstrap>
            <registered unixtime="1189696970">2007-09-13 15:22</registered>
            <type>subscriber</type>
        </user>
        ...
    </friends>
</lfm>

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getInfo

Get information about a user profile.
##### Example URLs

JSON: /2.0/?method=user.getinfo&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getinfo&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Optional) : The user to fetch info for. Defaults to the authenticated user.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<user>
    <id>1000002</id>
    <name>RJ</name>
    <realname>Richard Jones </realname>
    <url>http://www.last.fm/user/RJ</url>
    <image>http://userserve-ak.last.fm/serve/126/8270359.jpg</image>
    <country>UK</country>
    <age>27</age>
    <gender>m</gender>
    <subscriber>1</subscriber>
    <playcount>54189</playcount>
    <playlists>4</playlists>
    <bootstrap>0</bootstrap>
    <registered unixtime="1037793040">2002-11-20 11:50</registered>
</user>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getLovedTracks

Get the last 50 tracks loved by a user.
##### Example URLs

JSON: /2.0/?method=user.getlovedtracks&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getlovedtracks&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The user name to fetch the loved tracks for.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<lovedtracks user="RJ">
  <track>
    <name>The Glass Prison</name>
    <mbid/>
    <url>www.last.fm/music/Dream+Theater/_/The+Glass+Prison</url>
    <date uts="1216371514">18 Jul 2008, 08:58</date>
    <artist>
      <name>Dream Theater</name>
      <mbid>28503ab7-8bf2-4666-a7bd-2644bfc7cb1d</mbid>
      <url>http://www.last.fm/music/Dream+Theater</url>
    </artist>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </track>
  ...
</lovedtracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getPersonalTags

Get the user's personal tags

##### Example URLs

JSON: /2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR... (opens new window)
XML: /2.0/?method=user.getpersonaltags&user=rj&tag=rock&taggingtype=artist&api_key=YOUR... (opens new window)

##### Params

user (Required) : The user who performed the taggings.
tag (Required) : The tag you're interested in.
taggingtype[artist|album|track] (Required) : The type of items which have been tagged
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<taggings user="RJ" tag="rock" page="1" perPage="50" totalPages="1" total="11">
  <artists>
    <artist>
      <name>John Hammond</name>
      <mbid>d83e599c-2d5a-44ec-b727-587e1455b1b5</mbid>
      <url>http://www.last.fm/music/John+Hammond</url>
      <streamable>1</streamable>
      <image size="small">http://userserve-ak.last.fm/serve/34/255418.jpg</image>
      <image size="medium">http://userserve-ak.last.fm/serve/64/255418.jpg</image>
      <image size="large">http://userserve-ak.last.fm/serve/126/255418.jpg</image>
      <image size="extralarge">http://userserve-ak.last.fm/serve/252/255418.jpg</image>
      <image size="mega">http://userserve-ak.last.fm/serve/_/255418/John+Hammond.jpg</image>
    </artist>
  </artists>
</taggings>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getRecentTracks

Get a list of the recent tracks listened to by this user. Also includes the currently playing track with the nowplaying="true" attribute if the user is currently listening.

##### Example URLs

JSON: /2.0/?method=user.getrecenttracks&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getrecenttracks&user=rj&api_key=YOUR_API_KEY (opens new window)

##### Params

limit (Optional) : The number of results to fetch per page. Defaults to 50. Maximum is 200.
user (Required) : The last.fm username to fetch the recent tracks of.
page (Optional) : The page number to fetch. Defaults to first page.
from (Optional) : Beginning timestamp of a range - only display scrobbles after this time, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
extended (0|1) (Optional) : Includes extended data in each artist, and whether or not the user has loved each track
to (Optional) : End timestamp of a range - only display scrobbles before this time, in UNIX timestamp format (integer number of seconds since 00:00:00, January 1st 1970 UTC). This must be in the UTC time zone.
api_key (Required) : A Last.fm API key.

##### Auth

This service does not require authentication.

##### Sample Response

```
<recenttracks user="RJ" page="1" perPage="10" totalPages="3019">
  <track nowplaying="true">
    <artist mbid="2f9ecbed-27be-40e6-abca-6de49d50299e">Aretha Franklin</artist>
    <name>Sisters Are Doing It For Themselves</name>
    <mbid/>
    <album mbid=""/>
    <url>www.last.fm/music/Aretha+Franklin/_/Sisters+Are+Doing+It+For+Themselves</url>
    <date uts="1213031819">9 Jun 2008, 17:16</date>
    <streamable>1</streamable>
  </track>
  ...
</recenttracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getTopAlbums

Get the top albums listened to by a user. You can stipulate a time period. Sends the overall chart by default.

##### Example URLs

JSON: /2.0/?method=user.gettopalbums&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.gettopalbums&user=rj&api_key=YOUR_API_KEY (opens new window)

##### Params

user (Required) : The user name to fetch top albums for.
period (Optional) : overall | 7day | 1month | 3month | 6month | 12month - The time period over which to retrieve top albums for.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<topalbums user="RJ" type="overall">
  <album rank="1">
    <name>Images and Words</name>
    <playcount>174</playcount>
    <mbid>f20971f2-c8ad-4d26-91ab-730f6dedafb2</mbid>  
    <url>
      http://www.last.fm/music/Dream+Theater/Images+and+Words
    </url>
    <artist>
      <name>Dream Theater</name>
      <mbid>28503ab7-8bf2-4666-a7bd-2644bfc7cb1d</mbid>
      <url>http://www.last.fm/music/Dream+Theater</url>
    </artist>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </album>
</topalbums>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getTopArtists

Get the top artists listened to by a user. You can stipulate a time period. Sends the overall chart by default.
##### Example URLs

JSON: /2.0/?method=user.gettopartists&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.gettopartists&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The user name to fetch top artists for.
period (Optional) : overall | 7day | 1month | 3month | 6month | 12month - The time period over which to retrieve top artists for.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<topartists user="RJ" type="overall">
  <artist rank="1">
    <name>Dream Theater</name>
    <playcount>1337</playcount>
    <mbid>28503ab7-8bf2-4666-a7bd-2644bfc7cb1d</mbid>
    <url>http://www.last.fm/music/Dream+Theater</url>
    <streamable>1</streamable>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </artist>
  ...
</topartists>
```


##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getTopTags

Get the top tags used by this user.
##### Example URLs

JSON: /2.0/?method=user.gettoptags&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.gettoptags&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The user name
limit (Optional) : Limit the number of tags returned
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptags user="RJ">
  <tag>
    <name>rock</name>
    <count>12</count>
    <url>www.last.fm/tag/rock</url>
  </tag>
  ...
</toptags>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getTopTracks

Get the top tracks listened to by a user. You can stipulate a time period. Sends the overall chart by default.
##### Example URLs

JSON: /2.0/?method=user.gettoptracks&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.gettoptracks&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The user name to fetch top tracks for.
period (Optional) : overall | 7day | 1month | 3month | 6month | 12month - The time period over which to retrieve top tracks for.
limit (Optional) : The number of results to fetch per page. Defaults to 50.
page (Optional) : The page number to fetch. Defaults to first page.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<toptracks user="RJ" type="overall">  
  <track rank="1">
    <name>Learning to Live</name>
    <playcount>42</playcount>
    <mbid/>
    <url>
      http://www.last.fm/music/Dream+Theater/_/Learning+to+Live
    </url>
    <streamable fulltrack="0">1</streamable>
    <artist>
      <name>Dream Theater</name>
      <mbid>28503ab7-8bf2-4666-a7bd-2644bfc7cb1d</mbid>
      <url>http://www.last.fm/music/Dream+Theater</url>
    </artist>
    <image size="small">...</image>
    <image size="medium">...</image>
    <image size="large">...</image>
  </track>
  ...
</toptracks>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getWeeklyAlbumChart

Get an album chart for a user profile, for a given date range. If no date range is supplied, it will return the most recent album chart for this user.
##### Example URLs

JSON: /2.0/?method=user.getweeklyalbumchart&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getweeklyalbumchart&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The last.fm username to fetch the charts of.
from (Optional) : The date at which the chart should start from. See User.getChartsList for more.
to (Optional) : The date at which the chart should end on. See User.getChartsList for more.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<weeklyalbumchart user="RJ" from="1212321600" to="1212926400">
  <album rank="1">
    <artist mbid="80e577ba-841f-43ba-9f32-72e7c1692336">David Hudson</artist>
    <name>Bedarra</name>
    <mbid>dc30face-71db-413a-bcae-06accbd64aae</mbid>
    <playcount>10</playcount>
    <url>http://www.last.fm/music/David+Hudson+and+Friends/Bedarra</url>
  </album>
  ...
</weeklyalbumchart>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getWeeklyArtistChart

Get an artist chart for a user profile, for a given date range. If no date range is supplied, it will return the most recent artist chart for this user.
##### Example URLs

JSON: /2.0/?method=user.getweeklyartistchart&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getweeklyartistchart&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The last.fm username to fetch the charts of.
from (Optional) : The date at which the chart should start from. See User.getWeeklyChartList for more.
to (Optional) : The date at which the chart should end on. See User.getWeeklyChartList for more.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<weeklyartistchart user="RJ" from="1212321600" to="1212926400">
  <artist rank="1">
    <name>David Hudson and Friends</name>
    <mbid>80e577ba-841f-43ba-9f32-72e7c1692336</mbid>
    <playcount>18</playcount>
    <url>http://www.last.fm/music/David+Hudson+and+Friends</url>
  </artist>
  ...
</weeklyartistchart>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getWeeklyChartList

Get a list of available charts for this user, expressed as date ranges which can be sent to the chart services.
##### Example URLs

JSON: /2.0/?method=user.getweeklychartlist&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getweeklychartlist&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The last.fm username to fetch the charts list for.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<weeklychartlist user="RJ">
  <chart from="1108296002" to="1108900802"/>
  <chart from="1108900801" to="1109505601"/>
  ...
</weeklychartlist>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period

#### user.getWeeklyTrackChart

Get a track chart for a user profile, for a given date range. If no date range is supplied, it will return the most recent track chart for this user.
##### Example URLs

JSON: /2.0/?method=user.getweeklytrackchart&user=rj&api_key=YOUR_API_KEY&format=json (opens new window)
XML: /2.0/?method=user.getweeklytrackchart&user=rj&api_key=YOUR_API_KEY (opens new window)
##### Params

user (Required) : The last.fm username to fetch the charts of.
from (Optional) : The date at which the chart should start from. See User.getWeeklyChartList for more.
to (Optional) : The date at which the chart should end on. See User.getWeeklyChartList for more.
api_key (Required) : A Last.fm API key.
##### Auth

This service does not require authentication.
##### Sample Response

```
<weeklytrackchart user="joanofarctan" from="1212321600" to="1212926400">
  <track rank="1">
    <artist mbid="17b0d7f1-fad3-404e-87ae-874e6e158c3a">Dirk Leyers</artist>
    <name>Wellen</name>
    <mbid/>
    <playcount>3</playcount>
    <url>http://www.last.fm/music/Dirk+Leyers/_/Wellen</url>
  </track>
  ...
</weeklytrackchart>
```

##### Errors

- 2 : Invalid service - This service does not exist
- 3 : Invalid Method - No method with that name in this package
- 4 : Authentication Failed - You do not have permissions to access the service
- 5 : Invalid format - This service doesn't exist in that format
- 6 : Invalid parameters - Your request is missing a required parameter
- 7 : Invalid resource specified
- 8 : Operation failed - Something else went wrong
- 9 : Invalid session key - Please re-authenticate
- 10 : Invalid API key - You must be granted a valid key by last.fm
- 11 : Service Offline - This service is temporarily offline. Try again later.
- 13 : Invalid method signature supplied
- 16 : There was a temporary error processing your request. Please try again
- 26 : Suspended API key - Access for your account has been suspended, please contact Last.fm
- 29 : Rate limit exceeded - Your IP has made too many requests in a short period
