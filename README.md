# game-tracker

This project uses the Amazon Web Servirces Serverless Application Model (SAM)

**Project Structure**

```bash
.
├── Makefile                    <-- Make to automate build
├── README.md                   <-- This file
├── cmd                         <-- Individual lambda instance code
│   └── game                    <-- Lambda code
│         ├── game.go           <-- game API lambda handler
│         ├── game_test.go      <-- game API lambda tests
│         └── game_test_util.go <-- game API lambda test utilities
├── internal                    <-- Shared go code
│   └── data                    <-- Shared data structs
│         ├── data.go           <-- Common shared data structs
│         └── data_test.go      <-- Common shared data struct tests

```