package main

import(
  "fmt"
)

func main() {
  // Comic book variables
  var publisher, writer, artist, title, genre string
  var year, pageNumber, age int
  var grade, price float32

  // Define first comic book
  title = "Mr. GoToSleep"
  genre = "Mistery"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  age = 2020 - year
  price = float32(age * pageNumber) * grade / 100.0

  // Print out variables
  fmt.Println(
    title,
    "genre", genre,
    "written by", writer,
    "drawn by", artist,
    "published by", publisher,
    "on", year,
    "containing", pageNumber, "pages",
    "in condition", grade,
    "price", price,
  )

  // Define second comic book
  title = "Epic Vol. 1"
  genre = "SciFi"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0
  age = 2020 - year
  price = float32(age * pageNumber) * grade / 100.0

  // Print out variables
  fmt.Println(
    title,
    "genre", genre,
    "written by", writer,
    "drawn by", artist,
    "published by", publisher,
    "on", year,
    "containing", pageNumber, "pages",
    "in condition", grade,
    "price", price,
  )

  // Define third comic book
  title = "Ms. Y"
  genre = "Adventure"
  writer = "Gordon Ryan"
  artist = "Isobelle Leclair"
  publisher = "Astra Books"
  year = 2001
  pageNumber = 47
  grade = 8.5
  age = 2020 - year
  price = float32(age * pageNumber) * grade / 100.0

  // Print out variables
  fmt.Println(
    title,
    "genre", genre,
    "written by", writer,
    "drawn by", artist,
    "published by", publisher,
    "on", year,
    "containing", pageNumber, "pages",
    "in condition", grade,
    "price", price,
  )

}
