<template>
<div class="weather">
    <mt-header fixed title="天气预报">
		<div slot="right" @click="showSeven = !showSeven" v-html="btnText"></div>
	</mt-header>
    <div class="city-list">
		<div class="city-item" v-for="(item,index) in city" :key="index+'c'" >
		
            <p class="city-name">{{item.city}}</p>
			<div v-show="!showSeven">
				<p class="city-wea">{{item.data[0].wea}}</p>
            	<p class="city-temp">{{item.data[0].tem1 + " ~ " + item.data[0].tem2}}</p>
			</div>
            <div v-show="showSeven" v-for="(d,i) in item.data" :key="i+'x'" class="city-table">
				<table>
					<tbody>
						<tr>
							<td width="110px">{{d.day}}</td>
							<td>{{d.wea}}</td>
							<td width="110px">{{d.tem1 + " ~ " + d.tem2}}</td>
						</tr>
					</tbody>
				</table>
			</div>
        </div>
    </div>
</div>
</template>

<script>
import test from "../test.json";
import axios from "axios"
export default {
    name: "cityList",
    data() {
        return {
			city: [],
			showSeven:false
        };
    },
    created() {
		// this.city.push(...test.data)
		this.getWea()
    },
    computed: {
		btnText(){
			return this.showSeven ? "收回":"展开"
		}
	},
    methods: {
		getWea(){
			let v = this;
			axios.get("/api/weather").then( data => {
				v.city = []
				v.city.push(...data.data.data)
			})
		}
	}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->

<style lang="less">
.city-list {
    .city-item {
		box-sizing: border-box;
		float:left;
		width: 100%;
		padding: 10px 0;
		margin-top:10px;
		border-radius: 5px;
		font-size:16px;
        color: #fff;
        background: #24C6DC;
        /* fallback for old browsers */
        background: -webkit-linear-gradient(to left, #514A9D, #24C6DC);
        /* Chrome 10-25, Safari 5.1-6 */
        background: linear-gradient(to left, #514A9D, #24C6DC);
        /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

        p {
            margin: 0;
        }

        .city-name {
            font-size: 24px;
			font-weight: 600;
		}
		
		.city-table{
			table{
				width: 100%
			}
		}
    }
}
</style>
