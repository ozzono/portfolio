# CLI Technical Test

## Overview

You are only permitted to use library or language documentation or manual, no other sources of information may be used for solving the tasks.

## Background

Please implement a directory synchronization mechanism that will minimize the overall copy operation from the source directory to the target directory.  

### CLI USER INTERFACE  

#### Main Interface

|Command|Argument|Description|
|---|---|---|
|sync|-d|Destination directory|
||-s|Source, for this test the source is simply a directory in the file system|

## PROCESS

### On sync command

- The CLI application will start synchronizing the destination folder with the source folder, and the source folder should be regarded as the source of truth.

### The Source Directory File Structure

The contents of the source directory is completely arbitrary.

### OUTPUT

- On Success
  - Prints nothing
- On Error
  - Prints Err: \<Error Message\>

### CONSTRAINTS

- All of the arguments and their functions mentioned in the Main Interface of CLI USER INTERFACE section must be followed
- The application may not crash if an error happens, please display and return proper error code
- Consider only files, empty directories don't matter.
- You may not use any third party libraries.
- sort using any library are not allowed (use your own implementation of sorting)
- search using any library are not allowed (use your own implementation of searching)
- use of any library for getting list or info of file and/or folder are allowed
- For C++, you may not use the algorithm library from STL, the use of std::filesystem is permitted.
- You may use any language of your choice Note: if you apply for a language-specific position (e.g. C++), use that language.
- Create a fully object-oriented application
- Write unit / functional tests for the app.
- The app should be a CLI tool
- The app should display the result in the command prompt / terminal
- The app should produce the result file according to the input parameter
- Syncs are based on file name, do not need to check the file content.
