import React, { Component } from "react";
import { useNavigate } from "react-router-dom";

function HOCForRouteProps({ Component }) {
  const navigate = useNavigate();
  return <Component navigate={navigate} />;
}

export default HOCForRouteProps;
