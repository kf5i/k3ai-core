import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from "axios";

import PageLayout from "./components/PageLayout";
import Card from "./components/Card";

const App = () => {
  const [plugins, setPlugins] = useState([]);

  useEffect(() => {
    axios.get("/plugins").then(({ data }) => {
      setPlugins(data);
    });
  }, []);

  return (
    <PageLayout title="Plugins">
      {plugins.map(({ name: title, type: description }) => (
        <Card title={title} description={description} />
      ))}
    </PageLayout>
  );
};

ReactDOM.render(<App />, document.getElementById("k3ai-webui"));
