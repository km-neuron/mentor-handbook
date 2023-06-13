// src/features/todo/todoSlice.js

import { configureStore } from "@reduxjs/toolkit";

import todoReducer from "./slices/todoSlice";

const store = configureStore({
  reducer: {
    todos: todoReducer,
  },
});

// export store
export default store;