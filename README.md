<h1 align="center">
  🚀 🐹 My API Consumer
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>
<p align="center">
  This repository does various data retrievals for trusted sources automatically. It reads data from an input, makes a request, and records the result in a different file.
</p>

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.20 (recommended)

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/MyApiConsumer.git
```
3. Build up go project
```bash
go mod download
go get .
```
4. Make a **config.yaml** file in **./configs**
In **./configs/secure_config.yaml** has an example. Setup the data source clients and provide the location and name of input data.

The config files the following structure

| Father  |     Node     | Description                                                 | Possible values / Example                        |
|:-------:|:------------:|-------------------------------------------------------------|--------------------------------------------------|
|    -    |      io      | Describes Input parameters                                  | _Father node_                                    |
|   io    |   location   | Relative path where input can be found                      | ./file_exchange/                                 |
|   io    |  file_name   | Input filename (without extension)                          | input                                            |
|    -    |   clients    | Array with following elements, to describe each rest client | _Father node_                                    |
| clients |     name     | Name to invoke the client                                   | api-shazam                                       |
| clients |     type     | Resource type, see configs.go Resource_XXX constants        | GetRestApi / PostRestApi                         |
| clients |     path     | HTTP URL for client (without headers) *                     | https://shazam.p.rapidapi.com/shazam-events/list |
| clients |   headers    | Array with following elements, to describe each header      |                                                  |
| header  | :HEADER KEY: | Header name : header value                                  | 'X-RapidAPI-Key': 'SIGN-UP-FOR-KEY'              |

#### * Note about defining URIs:
If your URL contains path parameters, you may specify them with their names within curly brackets in the clients[n].path property. _For instance_: http://example.com/movies/{movie_name}

Make sure a column with these parameters exists in your input file. The next section contains further information on this input file.

5. Create an input file, in **config.yaml** file, you specified the property io.file_name. This property is a file where you specific the input data to be processed.

Csv files are one choice; the first row contains the name of each parameter, and each row contains one execution that defines all the parameters.

Example:

|  move_name   | year |  query  |
|:------------:|:----:|:-------:|
|    split     | 2016 | casting |
|  Inception   | 2010 | rating  |

The .csv file columns can be used like path parameters, query variables or both.

#### Json Body
You can use the first column on the input file to specify a JSON BODY. Just call the header as __"JSON_BODY"__ and each row the json in string format.

Example:

|    JSON_BODY    |  move_name  | year |  query  |
|:---------------:|:-----------:|:----:|:-------:|
| {"key":"value"} |    split    | 2016 | casting |
| {"key":"value"} |  Inception  | 2010 | rating  |

In this case, the first column will be treated like a Body in JSON format and no like a parameter. 

6. Run the API
```bash
go run cmd/cli/main.go
```
7. You'll see various errors since not enough parameters were provided; the correct form is as follows.
```bash
go run cmd/cli/main.go <DATA SOURCE>
```
This project use the following [flags](https://pkg.go.dev/flag):
- __inputter__: Input file type (_default "csv"_)
- __outputter__: Output file type (_default "json"_)

__Note__: No more clients are specified, but this code is extensible.

Thats why you don't need to specify the flags, but you can if you want to change the default values.

```bash
go run cmd/cli/main.go -inputter <INPUT FILE TYPE> -outputter <OUTPUT FILE TYPE> <DATA SOURCE>
```
For instance:
```bash
go run cmd/cli/main.go api_engine_beta
```
or 

```bash
go run cmd/cli/main.go -inputter csv -outputter json api_engine_beta 
```

A helper script is provided to run the API:

```bash
chmod +x runner.sh
```

```bash
./runner.sh
```