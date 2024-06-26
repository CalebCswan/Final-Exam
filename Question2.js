function reduceAndSortByFrequency(objects) {
    const arr = ["Baseball", "glove", "bat"];
    objects =
        { "value": "Baseball", "count": 100 },
        { "value": "glove", "count": 20 },
        { "value": "Baseball", "count": 15 },
        { "value": "glove", "count": 10 },
        { "value": "bat", "count": 25 },
        { "value": "bat", "count": 17 }

            arr.reduce 
            {((sorted) => index < sorted.length < sorted[index]);
             index++;}
            return sorted;
        
  }

  console.log(reduceAndSortByFrequency);