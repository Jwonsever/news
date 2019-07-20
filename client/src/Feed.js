import React, { useState, useEffect, useMemo } from "react";
import Parser from 'rss-parser'
import "./App.css";

function Feed() {
  const [feed, setFeed] = useState({});

  function isEmpty(obj) {
    if (Object.getOwnPropertyNames(obj).length > 0) return false;
    return true;
  }

  // Similar to componentDidMount and componentDidUpdate:
  useMemo(() => {
    let parser = new Parser();
     
    (async () => {
    
      let f = await parser.parseURL('http://localhost:8080');
      setFeed(f)
      console.log(feed.title);
     
      feed.items.forEach(item => {
        console.log(item.title + ':' + item.link)
      });
     
    })();

  }, [feed]);

  return (
    <div className="Feed">
    {feed.Name}
    </div>
  );
}

export default Feed;
