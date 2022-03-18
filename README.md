# <span style="color:#B4CDCD">arriva-cli</span>

A cli tool for the arriva api.
This is my first go project so the code might be a bit... poopie.

Build project:

```bash
go build -o bin/arriva-cli cmd/arriva-cli/main.go
```

## Usage:

```bash
./arriva-cli help   # Prints the help section (WIP)
./arriva-cli date   # Prints the current date (WIP)
./arriva-cli list   # Lists all stations
./arriva-cli "[start station]" "[end station]" [date] # Displays busses
```

## Examples:

Input:

```bash
./arriva-cli list | grep Mengeš
```

Output:

```
Mengeš Lek
Mengeš Lovec
Mengeš Pavovec
Mengeš Petrol
```

Input:

```bash
./arriva-cli "Mengeš Lovec" "Ljubljana Ap" 03-19
```

Output:

```
2022-03-19
*- Mengeš Lovec -- 05:20
│  ...
*- Ljubljana ap -- 05:54
Trajanje: 00:34

*- Mengeš Lovec -- 06:28
│  ...
*- Ljubljana ap -- 07:02
Trajanje: 00:34

*- Mengeš Lovec -- 07:22
│  ...
*- Ljubljana ap -- 07:50
Trajanje: 00:28

*- Mengeš Lovec -- 08:20
│  ...
*- Ljubljana ap -- 08:54
Trajanje: 00:34

*- Mengeš Lovec -- 09:20
│  ...
*- Ljubljana ap -- 09:54
Trajanje: 00:34

*- Mengeš Lovec -- 10:22
│  ...
*- Ljubljana ap -- 10:50
Trajanje: 00:28

*- Mengeš Lovec -- 11:20
│  ...
*- Ljubljana ap -- 11:54
Trajanje: 00:34

*- Mengeš Lovec -- 12:22
│  ...
*- Ljubljana ap -- 12:50
Trajanje: 00:28

*- Mengeš Lovec -- 13:20
│  ...
*- Ljubljana ap -- 13:54
Trajanje: 00:34

*- Mengeš Lovec -- 14:22
│  ...
*- Ljubljana ap -- 14:50
Trajanje: 00:28

*- Mengeš Lovec -- 15:20
│  ...
*- Ljubljana ap -- 15:54
Trajanje: 00:34

*- Mengeš Lovec -- 16:22
│  ...
*- Ljubljana ap -- 16:50
Trajanje: 00:28

*- Mengeš Lovec -- 17:20
│  ...
*- Ljubljana ap -- 17:54
Trajanje: 00:34

*- Mengeš Lovec -- 18:22
│  ...
*- Ljubljana ap -- 18:50
Trajanje: 00:28

*- Mengeš Lovec -- 19:20
│  ...
*- Ljubljana ap -- 19:54
Trajanje: 00:34

*- Mengeš Lovec -- 20:22
│  ...
*- Ljubljana ap -- 20:50
Trajanje: 00:28

*- Mengeš Lovec -- 20:22
│  ...
*- Ljubljana ap -- 20:50
Trajanje: 00:28

*- Mengeš Lovec -- 21:20
│  ...
*- Ljubljana ap -- 21:54
Trajanje: 00:34

*- Mengeš Lovec -- 22:20
│  ...
*- Ljubljana ap -- 22:54
Trajanje: 00:34
```
