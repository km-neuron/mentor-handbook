// src/redux/store.js

import { configureStore } from "@reduxjs/toolkit";

import todoReducer from "./slices/todoSlice";
import userReducer from "./slices/userSlice";

const store = configureStore({
  reducer: {
    todos: todoReducer,
    users: userReducer,
  },
  middleware: (getDefaultMiddleware) => getDefaultMiddleware(),
});

// export store
export default store;