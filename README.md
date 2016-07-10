### linktest

This is a script to crawl a website and search for broken links and imgs.

If the script finds no broken links or imgs, it passes.

If it finds one or more broken links or imgs, it fails, and outputs a list of the broken links (with both their source URLs and their destination URLs).

### Current status

This project is in early development.

So far, it can:
- Take a URL as an argument
- Find links on that page

It cannot yet:
- Dig into a second set of links
- Track broken links

As you continue development, pay close mind to:
- The concurrency model
- Timeouts
- Not checking the same link twice
- Differentiating between internal and external links
- Following relative links
