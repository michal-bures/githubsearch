# Github Exact Search
A small wrapper on top of the GitHub code search API that allows you to search for specific code fragments across all 
public repositories on GitHub.

Features:
 * Unlike the raw GitHub API, it sort of emulates ability to search for special characters and exact phrases (you can even use wildcards). 
 It does this by fetching first 500 results from the GitHub API and then manually filtering out any that don't actually contain 
 an exact match with the searched pattern.  
 * As a second step, it orderes the returned matches by how many watchers (stars) their repository has, starting with the most popular. 
 The reasoning is that more popular repositories are more likely to contain relevant code examples. 

It uses:

 * the built-in Go http server and HTML templating engine
 * the [go-github](https://github.com/google/go-github) GitHub api client 
