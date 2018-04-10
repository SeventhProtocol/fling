# Current Design Choices
Current, as in "today". Might change tomorrow. Or the day after. It's going to
be a lot of low-level framework design and coding first. I want to make sure
everything is on solid ground before implementing features. For that reason, the
request system will be the first main feature to be implemented. It's simple
enough to not steal all the focus, but requires most of the core functionality
(user management, networking, etc.).

I want to maintain many separate "cores" of Fling, called subsystems. Parts can
individually refactored or changed without affecting other portions. Because of
this, there's a lot of abstraction and communication between subsystems. As
things get going, I might combine or spin-off various functionalities. I really
can't stress enough that this is off-the-cuff design right now. But design and
talk is cheap, rather figure this out now then refactor the entirety of Fling
three months in.

## Language & Framework Choices
* __Language__: Golang
* __Web Framework__: Gin (most likely)
* __HTML Templating__: Either Jet or Hero. Leaning Jet.
* __Data Persistence__: SQLite
* __Configuration__: YAML (I hate configurations stored in SQL)

## Systems-Level Design
### Core
Responsible for data persistence, user management, authentication, &
permissions, scheduling tasks, and configuration.

### Notifications
Responsible for sending notifications on event triggers (new items fetched,
something requested, something went wrong, etc.).

### Fetcher
Responsible for searching and fetching for content. It does not manage requests,
or what's being tracked. A worker of sorts. 

### Queue
Responsible for maintaining requests and instructing the Fetcher to search for
things when it's time. Does not interact with anything external. That is handled
by other systems.

### Requests
Responsible for letting users add content to the Queue.

### Network
Responsible for all external network access. Must support proxies.

### File Management
Responsible for tracking what exists and what content goes where. 

### API
Responsible for a RESTful API that allows for configuration and direct
interaction with all other subsystems.

### Web Framework
A web management portal that interacts with the API subsystem to display
information and interact with Fling. Side note: Simultaneous work on the web
framework and a CLI will happen in the beginning. Having a CLI would be pretty
cool.

### Rosetta
Responsible for interacting with external resources. Works in tandem with the
Network subsystem to interact with torrent clients, data APIs (such as TMDB or
TVDB), and more.


<p align="center">
  <img src="https://i.imgur.com/UEdZpr4.png" alt="best girl"/>
</p>
