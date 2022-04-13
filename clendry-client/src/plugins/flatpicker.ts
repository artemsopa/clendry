import { App } from 'vue'
const flatPickr = require('vue-flatpickr-component');
import 'flatpickr/dist/flatpickr.css';

export function flatPickrInit(app: App<Element>) {
    app.use(flatPickr)
}
