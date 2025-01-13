import { configureStore, createSlice } from "@reduxjs/toolkit";

// Storage for message type
// For JavaScript, no type is needed, it's inferred
const initialState = {
  message: []
};

// App slice for managing messages
const appSlice = createSlice({
  name: "app",
  initialState,
  reducers: {
    sendMessage: (state, action) => {
      const newMessage = action.payload;
      state.message.push(newMessage);
    }
  }
});

// Export action and reducer
export const { sendMessage } = appSlice.actions;

const store = configureStore({
  reducer: {
    app: appSlice.reducer
  }
});

// Export store
export default store;
