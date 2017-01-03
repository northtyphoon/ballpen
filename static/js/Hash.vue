<template>
    <div id="hash">
        <p>{{ description }}</p>
        <p>
            <select v-model="algorithm">
                <option v-for="option in algorithmOptions" v-bind:value="option.value">
                    {{ option.text }}
                </option>
            </select>
            <input type="text" v-model="oriContent">
        </p>
        <p>{{ hashedContent }}</p>
    </div>
</template>

<script>
    import _ from 'lodash'
    import axios from 'axios'

    export default {
        data() {
            return {
                description: 'Hi, I\'m the hash ballpen. What can I hash for you?',
                algorithm: 'sha256',
                algorithmOptions: [
                    { text: 'SHA256', value: 'sha256' },
                    { text: 'MD5', value: 'md5' }
                ],
                oriContent: '',
                hashedContent: ''
            }
        },
        watch: {
            oriContent: function (newContent) {
                this.hashedContent = 'Waiting for you to stop typing...'
                this.caculateHashedContent()
            },
            algorithm: function (newContent) {
                if (this.oriContent) {
                    this.caculateHashedContent()
                }
            }
        },
        methods: {
            caculateHashedContent: _.debounce(
                function () {
                    var vm = this
                    if (vm.oriContent) {
                        vm.hashedContent = 'Hashing...'
                        axios.post('/api/hash', {
                            algorithm: vm.algorithm,
                            oriContent: vm.oriContent
                        }).then(function (response) {
                                vm.hashedContent = response.data.hashedContent
                            })
                            .catch(function (error) {
                                if (error.response) {
                                    // The request was made, but the server responded with a status code
                                    // that falls out of the range of 2xx
                                    vm.hashedContent = error.response.data.message
                                } else {
                                    // Something happened in setting up the request that triggered an Error
                                    vm.hashedContent = error.message;
                                }
                            })
                    }
                },
                // This is the number of milliseconds we wait for the
                // user to stop typing.
                500
            )
        }
    }
</script>

<style>
#hash input[type="text"] {
    text-align: center;
    width: 500px;
}
</style>