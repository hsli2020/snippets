https://gist.github.com/JeffreyWay/78f3bbac4564d2506916c1190b488aa7

// index.html
<!doctype html>
<html lang="en" class="h-full">
<head>
    <meta charset="UTF-8">
    <title>Episode 6: Bring it All Together</title>
    <script src="https://unpkg.com/vue@3"></script>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body class="h-full grid place-items-center bg-gray-800 text-white">
    <div id="app"></div>

    <script type="module">
        import App from './js/components/App.js';
        Vue.createApp(App).mount('#app');
    </script>
</body>
</html>

// js/components/App.js
import Assignments from "./Assignments.js";

export default {
    components: { Assignments },
    template: `<assignments></assignments>`,
}

// js/components/Assignment.js
export default {
    template: `
        <li>
            <label class="p-2 flex justify-between items-center">
                {{ assignment.name }}
                <input type="checkbox" v-model="assignment.complete" class="ml-3">
            </label>
        </li>
    `,

    props: { assignment: Object }
}

// js/components/AssignmentCreate.js
export default {
    template: `
        <form @submit.prevent="add">
            <div class="border border-gray-600 text-black">
                <input v-model="newAssignment" placeholder="New assignment..." class="p-2" />
                <button type="submit" class="bg-white p-2 border-l">Add</button>
            </div>
        </form>
    `,

    data() {
        return { newAssignment: '' }
    },

    methods: {
        add() {
            this.$emit('add', this.newAssignment);
            this.newAssignment = '';
        }
    }
}

// js/components/AssignmentList.js
import Assignment from "./Assignment.js";

export default {
    components: { Assignment },

    template: `
        <section v-show="assignments.length">
            <h2 class="font-bold mb-2">{{ title }}</h2>
            <ul class="border border-gray-600 divide-y divide-gray-600">
               <assignment
                    v-for="assignment in assignments"
                    :key="assignment.id"
                    :assignment="assignment"
                ></assignment>
            </ul>
        </section>
    `,

    props: {
        assignments: Array,
        title: String
    }
}

// js/components/Assignments.js
import AssignmentList from "./AssignmentList.js";
import AssignmentCreate from "./AssignmentCreate.js";

export default {
    components: { AssignmentList, AssignmentCreate },

    template: `
      <section class="space-y-6">
        <assignment-list :assignments="filters.inProgress" title="In Progress"></assignment-list>
        <assignment-list :assignments="filters.completed" title="Completed"></assignment-list>

        <assignment-create @add="add"></assignment-create>
      </section>
    `,

    data() {
        return {
            assignments: [
                { name: 'Finish project', complete: false, id: 1 },
                { name: 'Read Chapter 4', complete: false, id: 2 },
                { name: 'Turn in Homework', complete: false, id: 3 },
            ],
        }
    },

    computed: {
        filters() {
            return {
                inProgress: this.assignments.filter(assignment => ! assignment.complete),
                completed: this.assignments.filter(assignment => assignment.complete)
            };
        }
    },

    methods: {
        add(name) {
            this.assignments.push({
                name: name,
                completed: false,
                id: this.assignments.length + 1
            });
        }
    }
}
