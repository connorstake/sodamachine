import React, { Component } from "react";
import { useNavigate } from "react-router-dom";

function HOCForRouteProps({ Component }) {
  const navigate = useNavigate();
  console.log("HOC Props", navigate);
  return <Component navigate={navigate} />;
}

export default HOCForRouteProps;
