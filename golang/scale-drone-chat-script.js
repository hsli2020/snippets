// PS! Replace this with your own channel ID
const CLIENT_ID = 'YOUR_SCALEDRONE_ID';

// public room
const PUBLIC_ROOM_NAME = 'observable-room';

// array of connected memebers
let members = [];

// the session user
let me;

// keeping track of which room the user has selected
let selectedRoom = PUBLIC_ROOM_NAME;

// room name to messages map, this is used to store messages for displaying them
// at a later state
const roomMessages = {};

const drone = new Scaledrone(CLIENT_ID);

drone.on('open', error => {
  if (error) {
    return console.error(error);
  }
  // get JWT from the Go server for the clientID
  const formData = new FormData();
  formData.append('clientID', drone.clientId);
  fetch('/auth', {body: formData, method: 'POST'})
    .then(res => res.text())
    .then(jwt => drone.authenticate(jwt));
});

drone.on('authenticate', error => {
  if (error) {
    return console.error(error);
  }
  console.log('Successfully connected to Scaledrone');
  joinPublicRoom();
  joinPersonalRoom();
});

// Start subscribing to messages from the public room
function joinPublicRoom() {
  const publicRoom = drone.subscribe(PUBLIC_ROOM_NAME);
  publicRoom.on('open', error => {
    if (error) {
      return console.error(error);
    }
    console.log(`Successfully joined room ${PUBLIC_ROOM_NAME}`);
  });

  // Received array of members currently connected to the public room
  publicRoom.on('members', m => {
    members = m;
    me = members.find(m => m.id === drone.clientId);
    DOM.updateMembers();
  });

  // New member joined the public room
  publicRoom.on('member_join', member => {
    members.push(member);
    DOM.updateMembers();
  });

  // Member left public room (closed browser tab)
  publicRoom.on('member_leave', ({id}) => {
    const index = members.findIndex(member => member.id === id);
    members.splice(index, 1);
    DOM.updateMembers();
  });

  // Received public message
  publicRoom.on('message', message => {
    const {data, member} = message;
    if (member && member !== me) {
      addMessageToRoomArray(PUBLIC_ROOM_NAME, member, data);
      if (selectedRoom === PUBLIC_ROOM_NAME) {
        DOM.addMessageToList(data, member);
      }
    }
  });
}

// Start subscribing to messages from my private room (PMs to me)
function joinPersonalRoom() {
  const roomName = createPrivateRoomName(drone.clientId);
  const myRoom = drone.subscribe(roomName);
  myRoom.on('open', error => {
    if (error) {
      return console.error(error);
    }
    console.log(`Successfully joined room ${roomName}`);
  });

  myRoom.on('message', message => {
    const {data, clientId} = message;
    const member = members.find(m => m.id === clientId);
    if (member) {
      addMessageToRoomArray(createPrivateRoomName(member.id), member, data);
      if (selectedRoom === createPrivateRoomName(clientId)) {
        DOM.addMessageToList(data, member);
      }
    } else {
      /* Message is sent from golang using the REST API.
       * You can handle it like a regular message but it won't have a connection
       * session attached to it (this means no member argument)
       */
    }
  });
}

drone.on('close', event => {
  console.log('Connection was closed', event);
});

drone.on('error', error => {
  console.error(error);
});

function changeRoom(name, roomName) {
  selectedRoom = roomName;
  DOM.updateChatTitle(name);
  DOM.clearMessages();
  if (roomMessages[roomName]) {
    roomMessages[roomName].forEach(({data, member}) =>
      DOM.addMessageToList(data, member)
    );
  }
}

function createPrivateRoomName(clientId) {
  return `private-room-${clientId}`;
}

function addMessageToRoomArray(roomName, member, data) {
  console.log('add', roomName, member.id, data);
  roomMessages[roomName] = roomMessages[roomName] || [];
  roomMessages[roomName].push({member, data});
}

//------------- DOM Manipulation / Rendering the UI

const DOM = {
  elements: {
    me: document.querySelector('.me'),
    membersList: document.querySelector('.members-list'),
    messages: document.querySelector('.messages'),
    input: document.querySelector('.message-form__input'),
    form: document.querySelector('.message-form'),
    chatTitle: document.querySelector('.chat-title'),
    room: document.querySelector('.room'),
  },

  // Send message to Scaledrone and clear the input
  sendMessage() {
    const {input} = this.elements;
    const value = input.value;
    if (value === '') {
      return;
    }
    input.value = '';
    drone.publish({
      room: selectedRoom,
      message: value,
    });
    addMessageToRoomArray(selectedRoom, me, value);
    this.addMessageToList(value, me);
  },

  // Create DOM element with member name and color
  createMemberElement(member) {
    const { name, color } = member.authData;
    const el = document.createElement('div');
    el.appendChild(document.createTextNode(name));
    el.className = 'member';
    el.style.color = color;
    if (member !== me) {
      // Listen to user clicking on another user
      el.addEventListener('click', () =>
        changeRoom(member.authData.name, createPrivateRoomName(member.id))
      );
    }
    return el;
  },

  // Rerender the list of connected members
  updateMembers() {
    this.elements.me.innerHTML = '';
    this.elements.me.appendChild(this.createMemberElement(me));
    this.elements.membersList.innerHTML = '';
    members.filter(m => m !== me).forEach(member =>
      this.elements.membersList.appendChild(this.createMemberElement(member))
    );
  },

  // Create a DOM element for the message
  createMessageElement(text, member) {
    const el = document.createElement('div');
    el.appendChild(this.createMemberElement(member));
    el.appendChild(document.createTextNode(text));
    el.className = 'message';
    return el;
  },

  // Add message element to the messages container
  addMessageToList(text, member) {
    const el = this.elements.messages;
    const wasTop = el.scrollTop === el.scrollHeight - el.clientHeight;
    el.appendChild(this.createMessageElement(text, member));
    if (wasTop) {
      el.scrollTop = el.scrollHeight - el.clientHeight;
    }
  },

  updateChatTitle(roomName) {
    this.elements.chatTitle.innerText = roomName;
  },

  clearMessages() {
    this.elements.messages.innerHTML = '';
  },
};

// Listen to submitting the input form
DOM.elements.form.addEventListener('submit', () =>
  DOM.sendMessage()
);

// Listen to user clicking on the public room label
DOM.elements.room.addEventListener('click', () =>
  changeRoom('Public room', PUBLIC_ROOM_NAME)
);
