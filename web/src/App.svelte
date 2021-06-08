<script>
  import { onMount } from "svelte";
  //import { Tabs, Tab, TabList, TabPanel } from "svelte-tabs";
  // import { createForm } from "svelte-forms-lib";

  import TopAppBar, {
    Row,
    Section,
    Title,
    FixedAdjust,
    DenseFixedAdjust,
    ProminentFixedAdjust,
    ShortFixedAdjust
  } from "@smui/top-app-bar";

  import IconButton from "@smui/icon-button";
  import Devices from "./Devices.svelte";
  import Configuration from "./Configuration.svelte";
  import About from "./About.svelte";
  import Menu, { SelectionGroup, SelectionGroupIcon } from "@smui/menu";
  import List, {
    Item,
    Separator,
    Text,
    PrimaryText,
    SecondaryText,
    Graphic
  } from "@smui/list";

  // import * as yup from "yup";

  let prominent = false;
  let dense = false;
  let secondaryColor = false;

  let active = "Fermentation Devices";

  // const mqttEndpoint = "http://localhost:8080/v1/mqtt/";

  // let tilts = {};

  // onMount(async () => {
  //   const res1 = await fetch(mqttEndpoint);
  //   $form = await res1.json();
  //   // const res2 = await fetch(tiltEndpoint);
  //   // tilts = await res2.json();
  // });

  // async function doPut(body) {
  //   const res = await fetch(mqttEndpoint, {
  //     method: "PUT",
  //     body: body
  //   });
  // }

  // const { form, errors, state, handleChange, handleSubmit } = createForm({
  //   initialValues: {
  //     id: 1
  //     //   enabled: false,
  //     //   username: "",
  //     //   password: "",
  //     //   url: ""
  //   },
  //   validationSchema: yup.object().shape({
  //     enabled: yup.boolean().required(),
  //     url: yup
  //       .string()
  //       .matches(
  //         /^((mqtts?|tcp):)?\/\/(((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:)*@)?(((\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5])\.(\d|[1-9]\d|1\d\d|2[0-4]\d|25[0-5]))|((([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|\d|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.)+(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])*([a-z]|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])))\.?)(:\d*)?)(\/((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)+(\/(([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)*)*)?)?(\?((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)|[\uE000-\uF8FF]|\/|\?)*)?(\#((([a-z]|\d|-|\.|_|~|[\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF])|(%[\da-f]{2})|[!\$&'\(\)\*\+,;=]|:|@)|\/|\?)*)?$/i,
  //         "Not a valid URL"
  //       )
  //       .required(),
  //     username: yup.string().required(),
  //     password: yup.string().required()
  //   }),
  //   onSubmit: values => {
  //     doPut(JSON.stringify(values));
  //   }
  // });

  let menu;
  let clicked = "Fermentation Devices";
  //let dialog;
  let aboutComponent;

  function openAbout() {
    aboutComponent.openAbout();
  }
</script>

<style>
  :global(app, body, html) {
    display: block !important;
    height: auto !important;
    width: auto !important;
    position: static !important;
  }
  .top-app-bar-container {
    max-width: 480px;
    min-width: 480px;
    height: 320px;
    border: 1px solid rgba(0, 0, 0, 0.1);
    margin: 0 18px 18px 0;
  }
  .top-app-bar-container {
    overflow: auto;
    display: inline-block;
  }
  .flexy {
    display: flex;
    flex-wrap: wrap;
  }
</style>

<link
  rel="stylesheet"
  href="https://fonts.googleapis.com/icon?family=Material+Icons" />
<link
  rel="stylesheet"
  href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,600,700" />
<link
  rel="stylesheet"
  href="https://fonts.googleapis.com/css?family=Roboto+Mono" />

<div class="flexy">
  <div class="top-app-bar-container">
    <TopAppBar variant="static" {prominent} {dense} color="primary">
      <Row>
        <Section>
          <IconButton
            on:click={() => menu.setOpen(true)}
            class="material-icons">
            menu
          </IconButton>
          <Title>Brewery Monitor</Title>
        </Section>
      </Row>
    </TopAppBar>
  </div>
</div>

<Menu bind:this={menu}>
  <List>
    <Item on:SMUI:action={() => (clicked = 'Fermentation Devices')}>
      <Text>Fermentation Devices</Text>
    </Item>
    <Item on:SMUI:action={() => (clicked = 'Configuration')}>
      <Text>Configuration</Text>
    </Item>
    <Item on:click={() => openAbout()}>
      <Text>About</Text>
    </Item>
  </List>
</Menu>

{#if clicked == 'Configuration'}
  <Configuration />
{/if}

{#if clicked == 'Fermentation Devices'}
  <Devices />
{/if}

<!-- {#if clicked == 'About'} -->
<About bind:this={aboutComponent} />
<!-- {/if} -->

<!-- 
<Tabs>
  <TabList>
    <Tab>Devices</Tab>
    <Tab>MQTT</Tab>
    <Tab>Brewfather</Tab>
    <Tab>Grafana</Tab>
  </TabList>

  <TabPanel>
    <Devices />
  </TabPanel>

  <TabPanel>
    <form on:submit={handleSubmit}>

      <label for="enabled">enabled</label>
      <input
        type="checkbox"
        bind:checked={$form.enabled}
        id="enabled"
        name="enabled" />
      <br />

      <label for="url">url</label>
      <input
        id="url"
        name="url"
        on:change={handleChange}
        bind:value={$form.url} />
      {#if $errors.url}
        <small>{$errors.url}</small>
      {/if}

      <label for="username">name</label>
      <input
        id="username"
        name="username"
        on:change={handleChange}
        bind:value={$form.username} />
      {#if $errors.username}
        <small>{$errors.username}</small>
      {/if}

      <label for="password">password</label>
      <input
        id="password"
        name="password"
        on:change={handleChange}
        bind:value={$form.password} />
      {#if $errors.password}
        <small>{$errors.password}</small>
      {/if}

      <br />

      <button type="submit">Save</button>
    </form>
  </TabPanel>

  <TabPanel>Brewfather config here</TabPanel>

  <TabPanel />

</Tabs>
-->
