import React, { useState, useEffect } from "react";
import "./App.css";
import convert from "xml-js";

function Feed() {
  const [feed, setFeed] = useState({});

  function isEmpty(obj) {
    if (Object.getOwnPropertyNames(obj).length > 0) return false;
    return true;
  }

  // Similar to componentDidMount and componentDidUpdate:
  useEffect(() => {
    fetch("http://localhost:8080/feed")
      .then(response => response.text())
      .then(data => {
        data = convert.xml2js(data, {ignoreComment: true, alwaysChildren: true});
        let content = data.elements[0].elements[0].elements

        console.log("Debug log")
        console.log(content)

        // Parse this impossible to work with format into json.
        setFeed(content);

        if (!isEmpty(feed) && feed.Items !== undefined) {
          console.log(feed.title);
          feed.items.forEach(item => {
            console.log(item.title + ":" + item.link);
          });
        }
      });
  }, []);

  return <div className="Feed">{feed[0] ? feed[0].elements[0].text:""}</div>;
}

export default Feed;
