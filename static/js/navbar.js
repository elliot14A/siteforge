function navBarlinks() {
  const links = [
    {
      name: "Homepage",
      href: "/dashboard/content/home",
    },
    {
      name: "Project",
      href: "/dashboard/content/projects",
    },
    {
      name: "Gallery",
      href: "/dashboard/content/gallery",
    },
    {
      name: "Awards",
      href: "/dashboard/content/awards",
    },
    {
      name: "Certifications",
      href: "/dashboard/content/certifications",
    },
    {
      name: "About Us",
      href: "/dashboard/content/aboutus",
    },
  ];

  let active = "Homepage";

  return {
    links,
    active,
  };
}
