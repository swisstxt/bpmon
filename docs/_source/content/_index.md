---
title: "Home"
date: 2018-10-29T11:07:36+01:00
menu: "main"
weight: 1
---

## 'BPMON' what?

> BPMON is a command line tool that lets you compose Business Processes from the checks
in your monitoring system.

## Use Case

Often in IT a couple of tiny indicators we have in our monitoring system state whether an important 
service we provide is up or down. Let's assume we running an important web service. To make sure it
works we certainly have our monitoring system (eg. [ICINGA 2](https://icinga.com/products/icinga-2/))
and all required checks in place.

__But not every failed check is an incident!__ Given we have a tidy setup with a bunch of web servers,
a decent amount of database replicas etc. an outage of a single web server does not affect our service.

BPMON allows you to group checks into __Key Performance Indicators__ (KPI's). A KPI for our web service
could be *60% of our front end web servers must be healthy for the KPI to be healthy*. Neat!

A __Business Process__ (BP's) is a set of KPI's. Each KPI must be healthy for the BP to be healthy as well. 

## Features

BPMON is a very simple command line tool that...

* ... *reads* the states of services via ICINGA 2 (other monitoring systems could be supported on request),
* ... *evaluates* those states according to your definitions and rules, 
* ... triggers a command to perform actions such as *alarming*,
* ... feeds the check results to an InfluxDB in order to *keep track of the history* and
* ... (coming soon) provides a simple dashboard to access your availability data.

BPMON does not need to be run as a server. Run BPMON via Jenkins, Cron or
manually as needed.

## BPMON vs. ICINGA Web 2 Business Processes Module

You might be a bit confused since ICINGA already has a [Business Process Module](https://github.com/Icinga/icingaweb2-module-businessprocess).
So why even bother?!

Here's why:

1. BPMON currently only supports ICINGA 2, but can easily be extended to work with other monitoring solutions.
2. BPMON is completely configured via YAML files which can be managed via version control system.
3. Data generated by BPMON are stored in an efficient way (space wise and performance wise) in an Influx Database.
   you can easily access those data, integrate your alarming system and build dashboards for your customers.

If none of these points feel relevant to you or your business, you are probably better off with the ICINGA Module.
