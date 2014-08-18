ha_checker
==========

This is an application that given a set of backends to assigned to a port, it will respond to requests with an available backend.

Ideia
=====

The ideia is for any app to contact this app on a specific port and know the state of a backend.
This allows non-reliable applications to have a reliable backend choise. 

THis application will monitor single or groups of backends for availability and latency (future) and answer to requests
based on this metrics. The backend will be chosen based on the method selected (Random, Round-Robin, etc...)

Version
=======

0.0.1 - Just started
