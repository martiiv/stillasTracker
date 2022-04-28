import React from 'react'
import fetchModel from "./fetchData";
import {
    PROJECTS_URL_WITH_ID,
    PROJECTS_URL_WITH_SCAFFOLDING,
    SCAFFOLDING_URL,
    STORAGE_URL,
    WITH_SCAFFOLDING_URL
} from "./constantsFile";
import { useQuery, useQueryClient } from 'react-query'


export const GetDummyData = (dataName, url) => {
    const { isLoading, data} = useQuery(dataName, ()=>{
        return fetchModel(url)
    }, {
        refetchOnMount: false,
        refetchOnWindowFocus: false,
        refetchOnReconnect: false
    })

    return {isLoading, data}
}


 class AddData extends React.Component{
    constructor(pros) {
        super(pros);
    }

     async GetAllProjects() {
        let projects
        if (sessionStorage.getItem('allProjects') !== null) {
            projects = sessionStorage.getItem('allProjects')
        } else {
            console.log('From api')
            try {
                projects = fetchModel(PROJECTS_URL_WITH_SCAFFOLDING)
                projects = JSON.stringify(projects)
                sessionStorage.setItem('allProjects', projects)
            } catch (e) {
                console.log(e)
            }

        }

        return projects
    }

    async GetScaffoldingUnits(){
        let scaffolding
        if (sessionStorage.getItem('allScaffolding') !== null) {
            scaffolding = sessionStorage.getItem('allScaffolding')
        } else {
            console.log('From api')
            try {
                scaffolding = await fetchModel(SCAFFOLDING_URL)
                scaffolding = JSON.stringify(scaffolding)
                sessionStorage.setItem('allScaffolding', scaffolding)
            } catch (e) {
                console.log(e)
            }

        }

        return scaffolding
    }

     async GetStorageResult(){
         let storageResult
         if (sessionStorage.getItem('fromStorage') !== null) {
             storageResult = sessionStorage.getItem('fromStorage')
         } else {
             console.log('From api')
             try {
                 storageResult = await fetchModel(STORAGE_URL)
                 storageResult = JSON.stringify(storageResult)
                 sessionStorage.setItem('fromStorage', storageResult)
             } catch (e) {
                 console.log(e)
             }

         }
         return storageResult
     }

     async GetProject(id){
         let project
         if (sessionStorage.getItem('project') !== null) {
             project = sessionStorage.getItem('project')
         } else {
             console.log('From api')
             try {
                 project = await fetchModel(PROJECTS_URL_WITH_ID + id + WITH_SCAFFOLDING_URL )
                 project = JSON.stringify(project)
                 sessionStorage.setItem('project', project)
             } catch (e) {
                 console.log(e)
             }

         }

         return project
     }



 }

 export default AddData
